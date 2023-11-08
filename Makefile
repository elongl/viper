MAKEFILE_DIR_PATH = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
PROTO_PATH = protos/cmds.proto
AGENT_PATH = agent
CONTROLLER_PATH = controller
CONTROLLER_CLIENT_PATH = controller-client
BIN_PATH = $(MAKEFILE_DIR_PATH)bin

build: build_agent build_controller build_controller_client

build_agent:
	GOOS=windows GOARCH=amd64 go build -ldflags "-H=windowsgui -s -w" -o $(BIN_PATH)/agent.exe $(AGENT_PATH)/main.go
	go build -ldflags "-s -w" -o $(BIN_PATH)/agent $(AGENT_PATH)/main.go

build_agent_debug:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_PATH)/agent_debug.exe $(AGENT_PATH)/main.go
	go build -o $(BIN_PATH)/agent_debug $(AGENT_PATH)/main.go

build_controller:
	go build -o $(BIN_PATH)/controller $(CONTROLLER_PATH)/main.go

build_controller_client:
	pip install -r $(CONTROLLER_CLIENT_PATH)/requirements.txt

build_proto:
	protoc --go_out=$(AGENT_PATH)/protos $(PROTO_PATH)
	protoc --go_out=$(CONTROLLER_PATH)/protos --go-grpc_out=$(CONTROLLER_PATH)/protos $(PROTO_PATH)
	python -m grpc_tools.protoc -I. --python_out=$(CONTROLLER_CLIENT_PATH) --grpc_python_out=$(CONTROLLER_CLIENT_PATH) $(PROTO_PATH)

clean:
	rm -rf $(BIN_PATH)
