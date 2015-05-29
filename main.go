package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "rundeck-client"
	app.Usage = "Rundeck CLI tool"

	app.Commands = []cli.Command{
		{
			Name:    "project",
			Usage:   "Project commands",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List Projects: rundeck-client project list",
					Action: func(c *cli.Context) {
						cmd.ListProjects()
					},
				},
				{
					Name:  "executions",
					Usage: "List Executions: rundeck-client project executions projectid",
					Action: func(c *cli.Context) {
						var projectid string

						if len(c.Args()) != 1 {
							fmt.Printf("List Executions: rundeck-client project executions projectid\n")
							os.Exit(1)
						} else {
							projectid = c.Args()[0]
						}

						cmd.ListExecutions(projectid)
					},
				},
				{
					Name:  "history",
					Usage: "List History: rundeck-client project history projectid",
					Action: func(c *cli.Context) {
						var projectid string

						if len(c.Args()) != 1 {
							fmt.Printf("Usage: rundeck-client project history projectid\n")
							os.Exit(1)
						} else {
							projectid = c.Args()[0]
						}

						cmd.GetHistory(projectid)
					},
				},
                                {
                                        Name:  "execution-state",
                                        Usage: "Get Execution State: rundeck-client project executionid",
                                        Action: func(c *cli.Context) {
                                                var executionid string

                                                if len(c.Args()) != 1 {
                                                        fmt.Printf("Get Execution State: rundeck-client project executionid\n")
                                                        os.Exit(1)
                                                } else {
                                                        executionid = c.Args()[0]
                                                }

                                                cmd.GetExecutionstate(executionid)
                                        },
                                },

			},
		},
		{
			Name:    "job",
			Usage:   "job commands",
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
					Usage: "Get Job Options: rundeck-client job options jobid",
					Action: func(c *cli.Context) {
						var jobid string

						if len(c.Args()) != 1 {
							fmt.Printf("Get Job Options: rundeck-client job options jobid\n")
							os.Exit(1)
						} else {
							jobid = c.Args()[0]
						}

						cmd.GetRequiredJobOptions(jobid)
					},
				},
				{
					Name:  "run",
					Usage: "Run Job: rundeck-client job run jobid option1=option,option2=option",
					Action: func(c *cli.Context) {
						var jobid string
						var options string

						if len(c.Args()) <= 1 {
							fmt.Printf("Run Job: rundeck-client job run jobid option1=option,option2=option\n")
							os.Exit(1)
						} else {
							args := c.Args()
							jobid = args[0]
							l := len(args)
							s := 1
							for s < l {
								options = options + args[s]
								s++
							}
						}

						cmd.RunJob(jobid, options)
					},
				},

			},
		},
	}
	app.Run(os.Args)
}
