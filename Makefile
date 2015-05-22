all: clean test rundeck-client

test:
	#@go test rundeck-client -v

rundeck-client:
	@go get ./... 
	@go build 

clean:
	@rm -rf rundeck-client

.PHONY: all clean test rundeck-client
