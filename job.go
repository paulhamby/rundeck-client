package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
	"os"
	"strings"
)

var jobCommands = cli.Command{
	Name:  "job",
	Usage: "job commands",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "List Jobs: rundeck-client job list ",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				projectid := c.GlobalString("project")
				cmd.ListJobs(projectid)
			},
		},
		{
			Name:  "get",
			Usage: "Get Job: rundeck-client job get jobid",
			Action: func(c *cli.Context) {
				var jobid string

				if len(c.Args()) != 1 {
					fmt.Printf("Usage: rundeck-client job get jobid\n")
					os.Exit(1)
				} else {
					jobid = c.Args()[0]
				}

				cmd.GetJob(jobid)
			},
		},
		{
			Name:   "find",
			Usage:  "Find Job By Name: rundeck-client job find name",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				var jobid string

				if len(c.Args()) != 1 {
					fmt.Printf("Usage: rundeck-client job find\n")
					os.Exit(1)
				} else {
					jobid = c.Args()[0]
				}

				projectid := c.GlobalString("project")
				cmd.FindJobByName(jobid, projectid)
			},
		},
		{
			Name:   "options",
			Usage:  "Get Job Options: rundeck-client job options job project",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				var job string

				if len(c.Args()) != 1 {
					fmt.Printf("Get Job Options: rundeck-client job options job \n")
					os.Exit(1)
				} else {
					job = c.Args()[0]
				}

				projectid := c.GlobalString("project")
				cmd.GetJobOptions(job, projectid)
			},
		},
		{
			Name:   "run",
			Usage:  "Run Job: rundeck-client job run job option1=option,option2=option",
			Before: ensureProject,

			Action: func(c *cli.Context) {
				var options string
				args := c.Args()
				nbrArgsPassed := len(args)

				if nbrArgsPassed < 1 {
					fmt.Printf("Run Job: rundeck-client job run  job option1=option,option2=option\n")
					os.Exit(1)
				} else {
					s := 1
					for s < nbrArgsPassed {
						options = options + args[s] + " "
						s++
					}
					options = strings.TrimSpace(options)
				}

				projectid := c.GlobalString("project")
				cmd.RunJob(projectid, args[0], options)
			},
		},
	},
}
