package cmd

import (
	"fmt"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
	"strings"
)

func RunJob(projectid string, jobname string, options string) {
	var jobID string

	client := rundeck.NewClientFromEnv()

	jobByName, err := client.FindJobByName(jobname, projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		jobID = jobByName.ID
	}

	//input is quoted, comma-separated key/pairs: option1=option1,option2=option2
	//output should be : '-option1 option1 -option2 option2'
	var arguments string
	var loglevel string
	if options != "" {
		j := strings.Split(options, ",")
		for _, p := range j {
			s := strings.Split(p, "=")
			k, v := s[0], s[1]
			if k == "LogLevel" {
				loglevel = v
			} else {
				k = "-" + k
				arguments = arguments + " " + k + " " + v
			}
		}
	} else {
		arguments = ""
	}

	if loglevel == "" {
		loglevel = "INFO"
	}

	o := rundeck.RunOptions{LogLevel: loglevel, AsUser: "", Arguments: arguments}
	data, err := client.RunJob(jobID, o)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		var executionID string
		for _, d := range data.Executions {
			executionID = d.ID
		}
		GetExecutionstate(executionID, projectid)
		fmt.Printf("\nTo see the log from this execution, run 'rundeck-client execution output " + executionID + "'\n\n")
	}
}
