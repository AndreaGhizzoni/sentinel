package gitter

import (
	"bytes"
	"errors"
	"os/exec"
)

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
	//cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return "", err
	}

	//if len(stderr.String()) != 0 {
	//	return errors.New(stderr.String())
	//}

	return stdout.String(), nil
}
