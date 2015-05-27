package cmd

import (
	"fmt"
	"os"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
)

func RunJob(jobid string, options rundeck.RunOptions) {
	client := rundeck.NewClientFromEnv()
	data, err := client.RunJob(jobid, options)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID"})
		for _, d := range data.Executions {
			table.Append([]string{d.ID})
		}
		table.Render()
	}
}
