package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	. "github.com/paulhamby/go-rundeck/src/rundeck.v12"
	//"strconv"
	//"encoding/xml"
	"log"
	//"io/ioutil"
	//"strings"
)

func GetExecutionOutput(executionid string) {
	client := NewClientFromEnv()
	data, err := client.GetExecutionOutput(executionid)
	log.Printf("%+v", data.Entries)

	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Time",
			"Node",
			"Log",
		})
		for _, e := range data.Entries.Entry {
			fmt.Printf("%s\n", e.Log)
			//entry := formatOutputEntry(e.String())
			//entry = formatOutputEntry(e)
			//entry = entry + e.String()
			//fmt.Printf("%s\n", entry)
			table.Append([]string{ e.Time, e.Node, e.Log })
		}
		table.Render()
	}
}
