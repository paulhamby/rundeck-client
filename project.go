package main

import (
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
)

var projectCommands = cli.Command{
	Name:  "project",
	Usage: "Project commands",
	Subcommands: []cli.Command{
		{
			Name:  "list",
			Usage: "List Projects: rundeck-client project list",
			Action: func(c *cli.Context) {
				cmd.ListProjects()
			},
		},
		{
			Name:   "list-nodes",
			Usage:  "List Nodes: rundeck-client project list-nodes",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				projectid := c.GlobalString("project")
				cmd.ListNodes(projectid)
			},
		},
	},
}
