package main

import (
	"github.com/codegangsta/cli"
	"github.com/paulhamby/rundeck-client/cmd"
	"os"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-rundeck"
	app.Usage = "Rundeck CLI tool"

	app.Commands = []cli.Command{
		{
			Name:  "list-projects",
			Usage: "List Projects",
			Action: func(c *cli.Context) {
				cmd.ListProjects()
			},
		},
		{
			Name:  "list-jobs",
			Usage: "List Jobs: rundeck-client list-jobs projectid",
			Action: func(c *cli.Context) {
				var projectid string

				if len(c.Args()) != 1 {
					fmt.Printf("Usage: rundeck-client list-jobs projectid\n")
					os.Exit(1)
				} else {
					projectid = c.Args()[0]
				}

				cmd.ListJobs(projectid)
			},
		},
		{
			Name:  "list-executions",
			Usage: "List Executions: rundeck-client list-executions projectid",
			Action: func(c *cli.Context) {
				var projectid string

				if len(c.Args()) != 1 {
					fmt.Printf("Usage: rundeck-client list-executions projectid\n")
					os.Exit(1)
				} else {
					projectid = c.Args()[0]
				}

				cmd.ListExecutions(projectid)
			},
		},
		{
			Name:  "list-tokens",
			Usage: "List Tokens: rundeck-client list-tokens",
			Action: func(c *cli.Context) {
				cmd.ListTokens()
			},
		},
		{
			Name:  "get-history",
			Usage: "Get History: rundeck-client get-history projectid",
			Action: func(c *cli.Context) {
				var projectid string

				if len(c.Args()) != 1 {
					fmt.Printf("Usage: rundeck-client get-history projectid\n")
					os.Exit(1)
				} else {
					projectid = c.Args()[0]
				}

				cmd.GetHistory(projectid)
			},
		},
		{
			Name:  "get-job",
			Usage: "Get Job: rundeck-client get-job jobid",
			Action: func(c *cli.Context) {
				var jobid string

				if len(c.Args()) != 1 {
					fmt.Printf("Usage: rundeck-client get-job jobid\n")
					os.Exit(1)
				} else {
					jobid = c.Args()[0]
				}

				cmd.GetJob(jobid)
			},
		},
		{
			Name:  "find-job-by-name",
			Usage: "Find Job By Name: rundeck-client find-job-by-name jobid projectid",
			Action: func(c *cli.Context) {
				var jobid string
				var projectid string

				if len(c.Args()) != 2 {
					fmt.Printf("Usage: rundeck-client find-job-by-name jobid projectid\n")
					os.Exit(1)
				} else {
					jobid = c.Args()[0]
					projectid = c.Args()[1]
				}

				cmd.FindJobByName(jobid,projectid)
			},
		},
	}
	app.Run(os.Args)
}
