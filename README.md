# rundeck-client

A Rundeck client written in go. This is geared towards using Rundeck, as opposed to administering it.

Much of this is borrowed from github.com/lusis/go-rundeck. It is using rundeck.v12 as a library and re-implementing many of it's standalone binaries as functions so we can get one single command/entry point. Thanks lusis!

### Here's a way to get started

`go get github.com/paulhamby/rundeck-client`   
`export RUNDECK_URL=http://192.168.50.2:4440` (If you are using the anvils-demo vagrant image here: https://github.com/rundeck/anvils-demo)   
`export RUNDECK_TOKEN=asdfsdf` (Login to rundeck, go to your user profile and generate an API token if there isn't one there)   
`rundeck-client -h` (This will work if your PATH has $GOPATH/bin in it)   

### Example usage
```
$rundeck-client -h
NAME:
   rundeck-client - Rundeck CLI tool

USAGE:
   rundeck-client [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   project	Project commands
   execution	Execution commands
   job		job commands
   adhoc	Run adhoc commands and scripts
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --project, -p 		project
   --help, -h			show help
   --generate-bash-completion	
   --version, -v		print the version
   
```   
   

```
$rundeck-client -p anvils adhoc command 'echo Hello!; sleep 1; cat /etc/shadow' name=localhost
+----+---------+-----------+----------------------+------------+--------------------+
| ID | STATUS  | COMPLETED |        START         | STEP COUNT |     NODE STATE     |
+----+---------+-----------+----------------------+------------+--------------------+
| 60 | running | false     | 2015-06-07T01:35:17Z | 1          | localhost:RUNNING, |
+----+---------+-----------+----------------------+------------+--------------------+

To see the log from this execution, run 'rundeck-client execution output 60'

```

![Execution Output](/../screenshots/screenshots/execution-output.png?raw=true "Execution Output")


### To enable bash completion:

`sudo cp rundeck-client.bash_completion /etc/bash_completion.d/rundeck-client`

TODO
- add node list (done)  
- add node filters (done)   
- add tab/auto completion (done)   
- get the execution state of a job after executing it (done)   
- add ability to run adhoc scripts   
- add get all options function (instead of just required options) (done, thanks marybel!)   
- add ability to get the log from an execution (done)   
- add ability to run an adhoc command (done)   
- run job by name (done)   
- Put app commands in separate files for better layout
