PROTO_PATH = protos/cmds/cmds.proto
CONTROLLER_CLIENT_PATH = controller-client

build: build_agent build_agent_debug build_controller build_controller_client

build_agent:
	GOOS=windows go build -ldflags "-H=windowsgui -s -w" -o bin/agent.exe agent/main.go
	go build -ldflags "-s -w" -o bin/agent agent/main.go

build_agent_debug:
	GOOS=windows go build -o bin/agent_debug.exe agent/main.go
	go build -o bin/agent_debug agent/main.go

build_controller:
	go build -o bin/controller controller/main.go

build_controller_client:
	pip install -r $(CONTROLLER_CLIENT_PATH)/requirements.txt

build_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_PATH)
	python -m grpc_tools.protoc -I. --python_out=$(CONTROLLER_CLIENT_PATH) --grpc_python_out=$(CONTROLLER_CLIENT_PATH) $(PROTO_PATH)

clean:
	rm -rf bin
