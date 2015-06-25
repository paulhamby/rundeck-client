package cmd

import (
	"fmt"
	"os"
	//"strings"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)

func GetJobOptions(job string, projectid string) {
	var jobID string

	client := rundeck.NewClientFromEnv()

	jobByName, err := client.FindJobByName(job, projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	jobID = jobByName.ID

	data, err := client.GetOpts(jobID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name (* = Required)", "Default Value"})
	for n, v := range data {
		table.Append([]string{n, v})
	}
	table.Render()
}
