package cmd

import (
	"fmt"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
	"os"
)

func RunAdhoc(projectid string, exec string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.RunAdhoc(projectid, exec)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth(50)
		table.SetHeader([]string{"For job status run:"})
		table.Append([]string{"rundeck-client project execution-state " + data.ID + projectid})
		table.Render()
	}
}
