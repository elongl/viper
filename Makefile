build_agent:
	GOOS=windows go build -ldflags -H=windowsgui -o bin/agent.exe agent/main.go
	go build -o bin/agent agent/main.go

build_controller:
	go build -o bin/controller controller/main.go
