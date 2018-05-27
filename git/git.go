package git

import (
	"bytes"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	"log"
	"os/exec"
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
	baseHasBeenCreated, err := createFolderIfNotExists(workspace.Base)
	if err != nil {
		return "", err
	}
	if baseHasBeenCreated {
		g.log("Created\n")
	} else {
		g.log("Exists\n")
	}

	g.log("Unlocking Keys...\n")
	if err := g.unlockKeys(); err != nil {
		return "", err
	}

	green := color.New(color.FgHiGreen).SprintfFunc()
	for i := 0; i < len(workspace.Languages); i++ {
		var language = workspace.Languages[i]
		g.log(green("Processing Language: " + language.Name + "\n"))

		var languagePath = workspace.Base + "/" + language.Name
		_, err := createFolderIfNotExists(languagePath)
		if err != nil {
			return "", nil
		}

		for repoName, repoRemote := range language.Repositories {
			var repoPath = languagePath + "/" + repoName
			var err error
			var out string
			if folderNotExists(repoPath) {
				g.log("[cloning] " + repoRemote + " -> " + repoPath + " ... ")
				out, err = g.clone(repoRemote, repoPath)
			} else {
				g.log("[pulling] " + repoPath + " ... ")
				out, err = g.pull(repoPath)
			}
			g.log(out + "\n")

			if err != nil {
				return "", err
			}
		}
	}

	return "Finish.", nil
}

func (g *Gitter) unlockKeys() error {
	cmd := exec.Command("ssh-add", "-t", "90")
	_, err := g.runCommand(cmd)
	return err
}

func (g *Gitter) clone(repoRemote, localRepoName string) (string, error) {
	cmd := exec.Command("git", "clone", "-q", repoRemote, localRepoName)
	out, err := g.runCommand(cmd)
	if err != nil {
		return "", err
	}

	if len(out) == 0 {
		return "OK", nil
	}
	return out, nil
}

func (g *Gitter) pull(repoPath string) (string, error) {
	cmd := exec.Command("git", "-C", repoPath, "pull", "origin", "master")
	return g.runCommand(cmd)
}

func (g *Gitter) runCommand(cmd *exec.Cmd) (string, error) {
	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	//if len(stderr.String()) != 0 {
	//	return errors.New(stderr.String())
	//}

	return stdout.String(), nil
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
