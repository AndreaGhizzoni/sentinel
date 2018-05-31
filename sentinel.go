package main

import (
	"github.com/AndreaGhizzoni/sentinel/gitter"
	"github.com/AndreaGhizzoni/sentinel/net"
	"github.com/abiosoft/ishell"
	"github.com/urfave/cli"
	"os"
	"sort"
)

const (
	name    = "sentinel"
	version = "0.0.001"

	usage     = "TODO"
	usageText = "TODO"

	// commands
	shellCommand = "shell"
	shellUsage   = "TODO usage"
)

func readEnvironment() {
	gitter.AttachInput = os.Getenv("SSH_GRAPHIC_INPUT")
}

func main() {
	readEnvironment()

	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Usage = usage
	app.UsageText = usageText
	app.Authors = []cli.Author{
		{Name: "Andrea Ghizzoni", Email: "andrea.ghz@gmail.com"},
	}

	app.Commands = []cli.Command{
		{
			Name:   shellCommand,
			Usage:  shellUsage,
			Action: runAsShell,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

func runAsShell(c *cli.Context) {
	shell := ishell.New()

	shell.Println("")
	shell.Println("Welcome to Sentinel")

	var nmap = net.NewScanner("scan", "TODO help")
	var git = gitter.NewGitter("gitter", "TODO help")
	shell.AddCmd(nmap.GetIShellCommand())
	shell.AddCmd(git.GetIShellCommand())

	shell.Run()
}
