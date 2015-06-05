package cmd

import (
	"fmt"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)

func RunAdhoc(projectid string, exec string, node_filter string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.RunAdhoc(projectid, exec, node_filter)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		GetExecutionstate(data.ID, projectid)

	}
}
