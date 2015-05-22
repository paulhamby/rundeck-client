package main

import (
	"github.com/codegangsta/cli"
	"os"
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
				ListProjects()
			},
		},
		{
			Name:  "list-jobs",
			Usage: "List Jobs: rundeck-client list-jobs projectid",
			Action: func(c *cli.Context) {
				var projectid string

				if len(c.Args()) > 0 {
					projectid = c.Args()[0]
				}
				ListJobs(projectid)
			},
		},
	}
	app.Run(os.Args)
}
