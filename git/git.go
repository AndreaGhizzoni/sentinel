package git

import (
	"bytes"
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
)

type Gitter struct {
	name, help string
	cmd        *ishell.Cmd
}

func NewGitter(name, help string) *Gitter {
	return &Gitter{
		name: name,
		help: help,
	}
}

func (g *Gitter) Run(c *ishell.Context) (string, error) {
	workspace, err := NewWorkspace()
	if err != nil {
		return "", err
	}

	c.Println("Start checking workspace...")
	c.Printf("Base: %v .... ", workspace.Base)
	baseHasBeenCreated, err := createFolderIfNotExists(workspace.Base)
	if err != nil {
		return "", err
	}
	if baseHasBeenCreated {
		c.Println("Created")
	} else {
		c.Println("Exists")
	}

	if err := g.unlockKeys(); err != nil {
		return "", err
	}

	green := color.New(color.FgHiGreen).SprintfFunc()
	for i := 0; i < len(workspace.Languages); i++ {
		var language = workspace.Languages[i]
		c.Printf(green("Processing Language: %v\n", language.Name))

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
				c.Printf("[cloning] %v -> %v ... ", repoRemote, repoPath)
				out, err = g.clone(repoRemote, repoPath)
			} else {
				c.Printf("[pulling] %v ... ", repoPath)
				out, err = g.pull(repoPath)
			}
			c.Printf(out)

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

func (g *Gitter) clone(repoRemote, repoName string) (string, error) {
	cmd := exec.Command("git", "clone", "-q", repoRemote, repoName)
	out, err := g.runCommand(cmd)
	if err != nil {
		return "", err
	}

	if len(out) == 0 {
		return "OK\n", nil
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
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return "", err
	}

	//if len(stderr.String()) != 0 {
	//	return errors.New(stderr.String())
	//}

	return stdout.String(), nil
}

func (g *Gitter) run(c *ishell.Context) {
	res, err := g.Run(c)
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
