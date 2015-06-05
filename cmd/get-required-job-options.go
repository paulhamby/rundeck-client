package cmd

import (
	"fmt"
	"os"
	//"strings"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)

func GetRequiredJobOptions(job string, projectid string) {
        var jobID string

        client := rundeck.NewClientFromEnv()

        jobByName, err := client.FindJobByName(job, projectid)
        if err != nil {
                fmt.Printf("%s\n", err)
        } else {
                jobID = jobByName.ID
        }

	data, err := client.GetRequiredOpts(jobID)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Default Value"})
		for n, v := range data {
			table.Append([]string{n, v})
		}
		table.Render()
	}
}
