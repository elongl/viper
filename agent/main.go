package main

import (
	"log"
	"viper"
	"viper/agent/controller"
	"viper/agent/modules"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

func main() {
	controller := controller.Controller{Addr: viper.Conf.Agent.ControllerAddress}
	controller.Connect()
	for {
		cmdReq, err := controller.ReadCommandRequest()
		if err != nil {
			log.Printf("Failed to read command: %v", err)
			continue
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
