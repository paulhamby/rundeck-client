package main

import (
	"fmt"
	"os"
	"github.com/olekukonko/tablewriter"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
)

func ListTokens() {
	client := rundeck.NewClientFromEnv()
	data, err := client.GetTokens()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"User",
		})
		for _, d := range data.Tokens {
			table.Append([]string{
				d.ID,
				d.User,
			})
		}
		table.Render()
	}
}
