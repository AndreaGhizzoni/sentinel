package gitter

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

var AttachInput = "0"

func getProjectsCommand(cmd, repoRemote, localRepoName string) (*exec.Cmd, error) {
	command, err := getCommand(cmd)
	if err != nil {
		return nil, err
	}

	switch cmd {
	case "git":
		command.Args = append(command.Args, "clone", "-q", repoRemote, localRepoName)
	case "go":
		command.Args = append(command.Args, "-u", repoRemote)
	default:
		return nil, errors.New(cmd + ": Command Not found")
	}

	return command, nil
}

func getUpdateCommand(cmd, repoPath string) (*exec.Cmd, error) {
	command, err := getCommand(cmd)
	if err != nil {
		return nil, err
	}

	switch cmd {
	case "git":
		command.Args = append(command.Args, "-C", repoPath, "pull", "origin",
			"master")
	case "go":
		command.Args = append(command.Args, "-u", repoPath)
	default:
		return nil, errors.New(cmd + ": Command Not found")
	}

	return command, nil
}

func getCommand(cmd string) (*exec.Cmd, error) {
	switch cmd {
	case "git":
		return exec.Command("git"), nil
	case "go get":
		return exec.Command("go get"), nil
	case "ssh-add":
		return exec.Command("ssh-add"), nil
	default:
		return nil, errors.New(cmd + ": Command Not found")
	}
}

func runCommand(cmd *exec.Cmd) (string, error) {
	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if AttachInput == "0" {
		cmd.Stdin = os.Stdin
	}
	if err := cmd.Run(); err != nil {
		return "", err
	}

	//if len(stderr.String()) != 0 {
	//	return errors.New(stderr.String())
	//}

	return stdout.String(), nil
}
