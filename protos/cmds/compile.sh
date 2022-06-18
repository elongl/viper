#!/bin/sh

CONTROLLER_CLIENT_DIR_PATH="../../controller-client"

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cmds.proto
python -m grpc_tools.protoc -I. --python_out=$CONTROLLER_CLIENT_DIR_PATH --grpc_python_out=$CONTROLLER_CLIENT_DIR_PATH cmds.proto
