package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "rundeck-client"
	app.Usage = "Rundeck CLI tool"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
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
			},
		},
		{
			Name:  "execution",
			Usage: "Execution commands",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List Executions: rundeck-client execution list projectid",
					Action: func(c *cli.Context) {
						var projectid string

						if len(c.Args()) != 1 {
							fmt.Printf("List Executions: rundeck-client execution list projectid\n")
							os.Exit(1)
						} else {
							projectid = c.Args()[0]
						}

						cmd.ListExecutions(projectid)
					},
				},
				{
					Name:  "history",
					Usage: "List History: rundeck-client execution history projectid",
					Action: func(c *cli.Context) {
						var projectid string

						if len(c.Args()) != 1 {
							fmt.Printf("Usage: rundeck-client execution history projectid\n")
							os.Exit(1)
						} else {
							projectid = c.Args()[0]
						}

						cmd.GetHistory(projectid)
					},
				},
				{
					Name:  "state",
					Usage: "Get Execution State: rundeck-client execution state executionid projectid",
					Action: func(c *cli.Context) {
						var executionid string
						var projectid string

						if len(c.Args()) != 2 {
							fmt.Printf("Get Execution State: rundeck-client execution state executionid projectid\n")
							os.Exit(1)
						} else {
							executionid = c.Args()[0]
							projectid = c.Args()[1]
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
		},
		{
			Name:  "job",
			Usage: "job commands",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List Jobs: rundeck-client job list projectid",
					Action: func(c *cli.Context) {
						var projectid string

						if len(c.Args()) != 1 {
							fmt.Printf("List Jobs: rundeck-client job list projectid\n")
							os.Exit(1)
						} else {
							projectid = c.Args()[0]
						}

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
					Name:  "find",
					Usage: "Find Job By Name: rundeck-client job find name projectid",
					Action: func(c *cli.Context) {
						var jobid string
						var projectid string

						if len(c.Args()) != 2 {
							fmt.Printf("Usage: rundeck-client job find jobid projectid\n")
							os.Exit(1)
						} else {
							jobid = c.Args()[0]
							projectid = c.Args()[1]
						}

						cmd.FindJobByName(jobid, projectid)
					},
				},
				{
					Name:  "options",
					Usage: "Get Job Options: rundeck-client job options job project",
					Action: func(c *cli.Context) {
						var job string
						var project string

						if len(c.Args()) != 2 {
							fmt.Printf("Get Job Options: rundeck-client job options job project\n")
							os.Exit(1)
						} else {
							job = c.Args()[0]
							project = c.Args()[1]
						}

						cmd.GetRequiredJobOptions(job, project)
					},
				},
				{
					Name:  "run",
					Usage: "Run Job: rundeck-client job run projectid job option1=option,option2=option",
					Action: func(c *cli.Context) {
						var options string
						args := c.Args()
						nbrArgsPassed := len(args)

						if nbrArgsPassed < 2 {
							fmt.Printf("Run Job: rundeck-client job run projectid job option1=option,option2=option\n")
							os.Exit(1)
						} else {
							s := 2
							for s < nbrArgsPassed {
								options = options + args[s] + " "
								s++
							}
							options = strings.TrimSpace(options)
						}

						cmd.RunJob(args[0], args[1], options)
					},
				},
			},
		},
		{
			Name:  "adhoc",
			Usage: "Run adhoc commands and scripts",
			Subcommands: []cli.Command{
				{
					Name:  "command",
					Usage: "Run adhoc command: rundeck-client adhoc command projectid 'command' node-filter",
					Action: func(c *cli.Context) {
						var nodeFilter string
						args := c.Args()
						nbrArgsPassed := len(args)

						if nbrArgsPassed <= 2 {
							fmt.Printf("Run adhoc command: rundeck-client adhoc command projectid 'command' node-filter\n")
							os.Exit(1)
						} else {
							s := 2
							for s < nbrArgsPassed {
								nodeFilter = nodeFilter + args[s] + " "

								s++
							}
							nodeFilter = strings.TrimSpace(nodeFilter)
						}

						cmd.RunAdhoc(args[0], args[1], nodeFilter)
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
