package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
	"os"
)

var executionCommands = cli.Command{
	Name:  "execution",
	Usage: "Execution commands",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "List Executions: rundeck-client execution list ",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				projectid := c.GlobalString("project")
				cmd.ListExecutions(projectid)
			},
		},
		{
			Name:   "history",
			Usage:  "List History: rundeck-client execution history ",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				projectid := c.GlobalString("project")
				cmd.GetHistory(projectid)
			},
		},
		{
			Name:   "state",
			Usage:  "Get Execution State: rundeck-client execution state executionid ",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				projectid := c.GlobalString("project")
				var executionid string

				if len(c.Args()) != 1 {
					fmt.Printf("Get Execution State: rundeck-client execution state executionid \n")
					os.Exit(1)
				} else {
					executionid = c.Args()[0]
				}

				cmd.GetExecutionstate(executionid, projectid)
			},
		},
		{
			Name:  "output",
			Usage: "Get Execution Output: rundeck-client execution output executionid",
			Action: func(c *cli.Context) {
				var executionid string

				if len(c.Args()) < 1 {
					fmt.Printf("Get Execution Output: rundeck-client execution output executionid\n")
					os.Exit(1)
				} else {
					executionid = c.Args()[0]
				}

				cmd.GetExecutionOutput(executionid)
			},
		},
	},
}
