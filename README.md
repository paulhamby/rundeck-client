# rundeck-client

Most of this is borrowed from github.com/lusis/go-rundeck. It is using rundeck.v12 as a library and re-implementing his standalone binaries as functions so we can get one single command/entry point.

Hopefully lots more to be added soon.

TODO
- add node list
- add node filters (done)
- add tab/auto completion (done - kinda..needs documentation)
- get the execution state of a job after executing it
- add ability to run adhoc scripts
- add get all options function (instead of just required options)
- add ability to get the log from an execution (done)
- add ability to run an adhoc command (done)
- run job by name or id (done)
- use default variables when running a job
