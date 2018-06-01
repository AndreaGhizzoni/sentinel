package gitter

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

var AttachInput = "0"

func getProjectsCommand(cmd, repoRemote, repoName string) (*exec.Cmd, error) {
	command, err := getCommand(cmd)
	if err != nil {
		return nil, err
	}

	switch cmd {
	case "git":
		command.Args = append(command.Args, "clone", "-q", repoRemote, repoName)
	case "go":
		command.Args = append(command.Args, "get", "-u", repoRemote)
	default:
		return nil, errors.New(cmd + ": Command Not found")
	}

	return command, nil
}

func getUpdateCommand(cmd, repoRemote, repoPath string) (*exec.Cmd, error) {
	command, err := getCommand(cmd)
	if err != nil {
		return nil, err
	}

	switch cmd {
	case "git":
		command.Args = append(command.Args, "-C", repoPath, "pull", "origin", "master")
	case "go":
		command.Args = append(command.Args, "get", "-u", repoRemote)
	default:
		return nil, errors.New(cmd + ": Command Not found")
	}

	return command, nil
}

func getCommand(cmd string) (*exec.Cmd, error) {
	switch cmd {
	case "git":
		return exec.Command("git"), nil
	case "go":
		return exec.Command("go"), nil
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

	return stdout.String(), nil
}
