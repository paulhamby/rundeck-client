package cmd

import (
	"fmt"
	"os"
	//"strings"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)

func GetRequiredJobOptions(jobid string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.GetRequiredOpts(jobid)
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
