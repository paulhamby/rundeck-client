package cmd

import (
	"fmt"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ListProjects() {
	client := rundeck.NewClientFromEnv()
	data, err := client.ListProjects()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Name",
			"Description",
			"URL",
		})
		for _, d := range data.Projects {
			table.Append([]string{
				d.Name,
				d.Description,
				d.Url,
			})
			table.Render()
		}
	}

}
