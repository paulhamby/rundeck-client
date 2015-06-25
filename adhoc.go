package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
	"os"
	"strings"
)

var adhocCommands = cli.Command{
	Name:  "adhoc",
	Usage: "Run adhoc commands and scripts",
	Subcommands: []cli.Command{
		{
			Name:   "command",
			Usage:  "Run adhoc command: rundeck-client adhoc command 'command' node-filter",
			Before: ensureProject,
			Action: func(c *cli.Context) {
				var nodeFilter string
				args := c.Args()
				nbrArgsPassed := len(args)

				if nbrArgsPassed < 2 {
					fmt.Printf("Run adhoc command: rundeck-client adhoc command 'command' node-filter\n")
					os.Exit(1)
				} else {
					s := 1
					for s < nbrArgsPassed {
						nodeFilter = nodeFilter + args[s] + " "

						s++
					}
					nodeFilter = strings.TrimSpace(nodeFilter)
				}
				projectid := c.GlobalString("project")
				cmd.RunAdhoc(projectid, args[0], nodeFilter)
			},
		},
	},
}
