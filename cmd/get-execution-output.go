package cmd

import (
	"fmt"
	. "github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/mgutz/ansi"
)

func GetExecutionOutput(executionid string) {
	client := NewClientFromEnv()
	data, err := client.GetExecutionOutput(executionid)
	red := ansi.ColorCode("red")
	green := ansi.ColorCode("green")
	cyan := ansi.ColorCode("cyan")
	reset := ansi.ColorCode("reset")

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, e := range data.Entries.Entry {
		fmt.Printf("%s %s", e.Time, e.Node)
		if e.Level == "ERROR" {
			fmt.Println(red, e.Level, e.Log, reset)
		} else if e.Level == "VERBOSE" {
			fmt.Println(cyan, e.Level, e.Log, reset)
		} else {
			fmt.Println(green, e.Level, e.Log, reset)
		}
	}
}
