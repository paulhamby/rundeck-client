package cmd

import (
	"fmt"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func GetExecutionstate(executionid string, projectid string) {
	status, err := getExecutionStatus(executionid, projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	client := rundeck.NewClientFromEnv()
	var nodestate string
	data, err := client.GetExecutionState(executionid)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"Status",
		"Completed",
		"Start",
		"Step Count",
		"Node State",
	})
	for _, s := range data.Nodes {
		i := s.Name
		for _, e := range data.Steps {
			nodestate = nodestate + i + ":" + e.ExecutionState + ","
		}
	}
	table.Append([]string{
		strconv.FormatInt(data.ExecutionID, 10),
		status,
		strconv.FormatBool(data.Completed),
		data.StartTime,
		strconv.FormatInt(data.StepCount, 10),
		nodestate,
	})
	table.Render()
}

func getExecutionStatus(executionid string, projectid string) (string, error) {
	client := rundeck.NewClientFromEnv()
	options := make(map[string]string)
	options["max"] = "200"
	options["project"] = projectid
	var status string
	data, err := client.ListExecutions(projectid, options)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		for _, d := range data.Executions {
			if d.ID == executionid {
				status = d.Status
			}
		}
	}
	return status, err
}
