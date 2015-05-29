package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)

func GetExecutionstate(executionId string) {
	client := rundeck.NewClientFromEnv()
	var nodestate string
	data, err := client.GetExecutionState(executionId)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"Success",
			"Completed",
			"Start",
			"Step Count",
			"Node State",
		})
		for _, s := range data.Nodes {
			i := s.Name
			for _, e := range data.Steps {
				nodestate = nodestate + i + ":" +  e.ExecutionState + ","
			}
		}
		table.Append([]string{
			strconv.FormatInt(data.ExecutionID, 10),
			strconv.FormatBool(data.Success),
			strconv.FormatBool(data.Completed),
			data.StartTime,
			strconv.FormatInt(data.StepCount, 10),
			nodestate,
		})
		table.Render()
	}
}
