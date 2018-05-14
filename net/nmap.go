package net

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/abiosoft/ishell"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Scanner struct {
	name, help string
	cmd        *ishell.Cmd
}

func NewScanner(name, help string) *Scanner {
	return &Scanner{
		name: name,
		help: help,
	}
}

func (w *Scanner) Run() (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("nmap", "-sP", "192.168.2.0/24", "-oX", "out.xml")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	if len(stderr.String()) != 0 {
		return "", errors.New(stderr.String())
	}

	result, err := w.processStdOut("out.xml")
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func (w *Scanner) processStdOut(outFile string) (*Result, error) {
	xmlFile, err := os.Open("out.xml")
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var res = &Result{}
	err = xml.Unmarshal(byteValue, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (w *Scanner) run(c *ishell.Context) {
	c.ProgressBar().Indeterminate(true)
	c.Println("Scanning...")
	c.ProgressBar().Start()
	out, err := w.Run()
	c.ProgressBar().Stop()

	if err != nil {
		log.Printf("Error %v", err)
	}
	c.Println(out)
}

func (w *Scanner) GetIShellCommand() *ishell.Cmd {
	if w.cmd == nil {
		w.cmd = &ishell.Cmd{
			Name: w.name,
			Help: w.help,
			Func: w.run,
		}
	}
	return w.cmd
}
