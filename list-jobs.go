package main

import (
	"fmt"
	"os"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
)

func ListJobs(projectid string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.ListJobs(projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description", "Group", "Project"})
		for _, d := range data.Jobs {
			table.Append([]string{d.ID, d.Name, d.Description, d.Group, d.Project})
		}
		table.Render()
	}
}
