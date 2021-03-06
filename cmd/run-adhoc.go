package cmd

import (
	"fmt"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
)

func RunAdhoc(projectid string, exec string, node_filter string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.RunAdhoc(projectid, exec, node_filter)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	GetExecutionstate(data.ID, projectid)
	fmt.Printf("\nTo see the log from this execution, run 'rundeck-client execution output %s'\n\n", data.ID)
}
