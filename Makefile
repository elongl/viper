build: build_agent build_agent_debug build_controller

build_agent:
	GOOS=windows go build -ldflags "-H=windowsgui -s -w" -o bin/agent.exe agent/main.go
	go build -ldflags "-s -w" -o bin/agent agent/main.go

build_agent_debug:
	GOOS=windows go build -o bin/agent_debug.exe agent/main.go
	go build -o bin/agent_debug agent/main.go

build_controller:
	go build -o bin/controller controller/main.go

clean:
	rm -rf bin
