package gitter

import (
	"bytes"
	"os/exec"
)

var sshAddCommand = exec.Command("ssh-add")
var sshAddArgs = []string{"-t", "90"}

var gitCommand = exec.Command("git")
var goCommand = exec.Command("go get")

func runCommand(cmd *exec.Cmd) (string, error) {
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
