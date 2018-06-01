package gitter

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	"log"
	"os"
)

type Gitter struct {
	name, help string
	cmd        *ishell.Cmd
	context    *ishell.Context
}

func NewGitter(name, help string) *Gitter {
	return &Gitter{
		name: name,
		help: help,
	}
}

func (g *Gitter) log(msg string) {
	if g.context != nil {
		g.context.Print(msg)
	} else {
		fmt.Print(msg)
	}
}

func (g *Gitter) Run() (string, error) {
	workspace, err := NewWorkspace()
	if err != nil {
		return "", err
	}

	g.log("Start checking workspace...\n")
	g.log("Base: " + workspace.Base + " .... ")
	if baseHasBeenCreated, err := workspace.BuildBaseFolderStructure(); err != nil {
		return "", err
	} else {
		if baseHasBeenCreated {
			g.log("Created\n")
		} else {
			g.log("Exists\n")
		}
	}

	g.log("Unlocking Keys...\n")
	if err := g.unlockKeys(); err != nil {
		return "", err
	}

	green := color.New(color.FgHiGreen).SprintfFunc()
	for _, language := range workspace.Languages {
		g.log(green("Processing Language: " + language.Name + "\n"))

		language.BuildFolderStructure(workspace.Base)

		for repoName, repoRemote := range language.Repositories {
			var repoPath = language.ProjectsFolder + "/" + repoName
			var err error
			var out string
			if folderNotExists(repoPath) {
				g.log("[cloning] " + repoRemote + " -> " + repoPath + " ... ")
				out, err = g.getProject(language.Command, repoRemote, repoPath, language.ProjectsFolder)
			} else {
				g.log("[pulling] " + repoPath + " ... ")
				out, err = g.updateProject(language.Command, repoRemote, repoPath, language.ProjectsFolder)
			}
			g.log(out)

			if err != nil {
				return "", err
			}
		}
	}

	return "Finish.", nil
}

func (g *Gitter) unlockKeys() error {
	cmd, _ := getCommand("ssh-add")
	cmd.Args = append(cmd.Args, "-t", "90")
	_, err := runCommand(cmd)
	return err
}

func (g *Gitter) getProject(cmd, repoRemote, repoName, projectsFolder string) (string, error) {
	command, err := getProjectsCommand(cmd, repoRemote, repoName)
	if err != nil {
		return "", err
	}

	os.Setenv("GOPATH", projectsFolder)
	out, err := runCommand(command)
	if err != nil {
		return "", nil
	}

	if len(out) == 0 {
		return "OK\n", nil
	}

	return out, nil
}

func (g *Gitter) updateProject(cmd, repoRemote, repoPath, projectsFolder string) (string, error) {
	command, err := getUpdateCommand(cmd, repoRemote, repoPath)
	if err != nil {
		return "", nil
	}

	os.Setenv("GOPATH", projectsFolder)
	out, err := runCommand(command)
	if err != nil {
		return "", nil
	}

	if len(out) == 0 {
		return "OK\n", nil
	}

	return out, nil
}

func (g *Gitter) run(c *ishell.Context) {
	g.context = c
	res, err := g.Run()
	if err != nil {
		log.Printf("Error %v", err)
	}
	c.Println(res)
}

func (g *Gitter) GetIShellCommand() *ishell.Cmd {
	if g.cmd == nil {
		g.cmd = &ishell.Cmd{
			Name: g.name,
			Help: g.help,
			Func: g.run,
		}
	}
	return g.cmd
}
