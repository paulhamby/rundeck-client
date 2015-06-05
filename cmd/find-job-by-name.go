package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)

func FindJobByName(jobid string, projectid string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.FindJobByName(jobid, projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		scope := *data
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description", "Group", "Steps", "Node Filters"})
		var steps []string
		var nodefilters []string
		for _, d := range scope.Sequence.Steps {
			var stepDescription string
			if d.Description == "" {
				if d.JobRef != nil {
					stepDescription = d.JobRef.Name
				} else if d.Exec != nil {
					stepDescription = *d.Exec
				}
			} else {
				stepDescription = d.Description
			}
			steps = append(steps, stepDescription)
		}
		for _, n := range scope.NodeFilters.Filter {
			nodefilters = append(nodefilters, n)
		}
		table.Append([]string{scope.ID, scope.Name, scope.Description, scope.Group, strings.Join(steps, "\n"), strings.Join(nodefilters, "\n")})
		table.Render()
	}
}
