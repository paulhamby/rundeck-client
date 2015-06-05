package cmd

import (
	"fmt"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
	"os"
	"strings"
)

func RunJob(jobid string, options string) {
	//input is quoted, comma-separated key/pairs: option1=option1,option2=option2
	//output should be : '-option1 option1 -option2 option2'
	var i string
	j := strings.Split(options, ",")
	for _, p := range j {
		fmt.Print(p + "\n")
		s := strings.Split(p, "=")
		k, v := s[0], s[1]
		k = "-" + k
		i = i + " " + k + " " + v
		fmt.Print(i + "\n")
	}
	o := rundeck.RunOptions{LogLevel: "DEBUG", AsUser: "", Arguments: i}
	client := rundeck.NewClientFromEnv()
	data, err := client.RunJob(jobid, o)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth(50)
		table.SetHeader([]string{"For job status run:"})
		for _, d := range data.Executions {
			table.Append([]string{"rundeck-client project execution-state " + d.ID})
		}
		table.Render()
	}
}
