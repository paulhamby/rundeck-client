package cmd

import (
	"fmt"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"os"
)

func GetJobOptions(job string, projectid string) {
	var jobID string
	options := make(map[string]string)

	client := rundeck.NewClientFromEnv()

	jobByName, err := client.FindJobByName(job, projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	jobID = jobByName.ID

	data, err := client.GetJob(jobID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name (* = Required)", "Default Value"})
	if data.Job.Context.Options != nil {
		for _, option := range *data.Job.Context.Options {
			var optionRequirement string
			if option.Required {
				optionRequirement = "* "
			} else {
				optionRequirement = "  "
			}

			var optionValue string
			if option.DefaultValue == "" {
				optionValue = "<no default>"
			} else {
				optionValue = option.DefaultValue
			}

			options[optionRequirement+option.Name] = optionValue
		}
	}

	for n, v := range options {
		table.Append([]string{n, v})
	}
	table.Render()
}
