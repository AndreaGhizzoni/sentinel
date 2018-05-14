package main

import (
	"github.com/AndreaGhizzoni/sentinel/net"
	"github.com/abiosoft/ishell"
)

func main() {
	shell := ishell.New()

	shell.Println("")
	shell.Println("Welcome to Sentinel")

	var nmap = net.NewScanner("scan", "help")
	shell.AddCmd(nmap.GetIShellCommand())

	shell.Run()
}
