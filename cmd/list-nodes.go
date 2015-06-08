package cmd

import (
	"fmt"
	"os"
	//"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"github.com/paulhamby/go-rundeck/src/rundeck.v12"
)
/*
  <node name="app1.anvils.com" description="A app server node." tags="anvils, app" hostname="localhost" osArch="x86_64" osFamily="unix" osName="Linux" osVersion="2.6.32-279.el6.x86_64" username="app1" ssh-key-storage-path="/keys/acme/anvils/app1/id_rsa">
    <attribute name="anvils:location" value="US-East"/>
    <attribute name="anvils:server-pool-id" value="1"/>
    <attribute name="anvils:server-pool" value="app"/>
    <attribute name="anvils:customer" value="acme.com"/>
  </node>
*/
func ListNodes(projectid string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.ListNodes(projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description", "Tags", "Hostname", "OsArch", "OsName", "OSVersion"})
		for _, d := range data.Nodes {
			table.Append([]string{d.Name, d.Description, d.Tags, d.Hostname, d.OsArch, d.OsName, d.OsVersion})
		}
		table.Render()
	}
}
