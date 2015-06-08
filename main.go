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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "project, p",
			Usage: "project",
		},
	}

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
				{
					Name:  "list-nodes",
					Usage: "List Nodes: rundeck-client project list-nodes",
					Before: func(c *cli.Context) error {
						ensureProject(c.GlobalString("project"))
						return nil
					},
					Action: func(c *cli.Context) {
						projectid := c.GlobalString("project")
						cmd.ListNodes(projectid)
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
					Usage: "List Executions: rundeck-client execution list ",
					Before: func(c *cli.Context) error {
						ensureProject(c.GlobalString("project"))
						return nil
					},
					Action: func(c *cli.Context) {
						projectid := c.GlobalString("project")
						cmd.ListExecutions(projectid)
					},
				},
				{
					Name:  "history",
					Usage: "List History: rundeck-client execution history ",
					Before: func(c *cli.Context) error {
						ensureProject(c.GlobalString("project"))
						return nil
					},
					Action: func(c *cli.Context) {
						projectid := c.GlobalString("project")
						cmd.GetHistory(projectid)
					},
				},
				{
					Name:  "state",
					Usage: "Get Execution State: rundeck-client execution state executionid ",
					Before: func(c *cli.Context) error {
						ensureProject(c.GlobalString("project"))
						return nil
					},

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
		},
		{
			Name:  "job",
			Usage: "job commands",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "List Jobs: rundeck-client job list ",
					Before: func(c *cli.Context) error {
						ensureProject(c.GlobalString("project"))
						return nil
					},
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
					Name:  "find",
					Usage: "Find Job By Name: rundeck-client job find name",
                                        Before: func(c *cli.Context) error {
                                                ensureProject(c.GlobalString("project"))
                                                return nil
                                        },
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
					Name:  "options",
					Usage: "Get Job Options: rundeck-client job options job project",
                                        Before: func(c *cli.Context) error {
                                                ensureProject(c.GlobalString("project"))
                                                return nil
                                        },
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
					Name:  "run",
					Usage: "Run Job: rundeck-client job run job option1=option,option2=option",
                                        Before: func(c *cli.Context) error {
                                                ensureProject(c.GlobalString("project"))
                                                return nil
                                        },
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
		},
		{
			Name:  "adhoc",
			Usage: "Run adhoc commands and scripts",
			Subcommands: []cli.Command{
				{
					Name:  "command",
					Usage: "Run adhoc command: rundeck-client adhoc command 'command' node-filter",
                                        Before: func(c *cli.Context) error {
                                                ensureProject(c.GlobalString("project"))
                                                return nil
                                        },
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
		},
	}
	app.Run(os.Args)
}

func ensureProject(project string) {
	if project == "" {
		fmt.Printf("Project flag must be set. Example: rundeck-client --project anvils\n\n")
		os.Exit(1)
	} else {
		return
	}
}
