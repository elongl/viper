MAKEFILE_DIR_PATH = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
PROTO_PATH = protos/cmds.proto
AGENT_PATH = agent
CONTROLLER_PATH = controller
CONTROLLER_CLIENT_PATH = controller-client
BIN_PATH = $(MAKEFILE_DIR_PATH)bin

build: build_agent build_agent_debug build_controller build_controller_client

build_agent:
	GOOS=windows go build -ldflags "-H=windowsgui -s -w" -o bin/agent.exe agent/main.go
	go build -ldflags "-s -w" -o bin/agent agent/main.go

build_agent_debug:
	GOOS=windows go build -o $(BIN_PATH)/agent_debug.exe main.go
	go build -o $(BIN_PATH)/agent_debug main.go

build_controller:
	go build -o bin/controller controller/main.go

build_controller_client:
	pip install -r $(CONTROLLER_CLIENT_PATH)/requirements.txt

build_proto:
	protoc --go_out=$(AGENT_PATH)/protos $(PROTO_PATH)
	protoc --go_out=$(CONTROLLER_PATH)/protos --go-grpc_out=$(CONTROLLER_PATH)/protos $(PROTO_PATH)
	python -m grpc_tools.protoc -I. --python_out=$(CONTROLLER_CLIENT_PATH) --grpc_python_out=$(CONTROLLER_CLIENT_PATH) $(PROTO_PATH)

install_proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	pip install grpcio-tools

clean:
	rm -rf bin
