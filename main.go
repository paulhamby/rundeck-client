package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
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
		projectCommands,
		executionCommands,
		jobCommands,
		adhocCommands,
	}
	app.Run(os.Args)
}

var ensureProject = func(c *cli.Context) error {
	project := c.GlobalString("project")
	if project == "" {
		fmt.Printf("Project flag must be set. Example: rundeck-client --project anvils\n\n")
		os.Exit(1)
	}
	return nil
}
