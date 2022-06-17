package main

import (
	"flag"
	"log"
	"viper/agent/controller"
	"viper/agent/modules"
	pb "viper/protos/cmds"
)

var (
	addr = flag.String("addr", "localhost:50051", "The address to connect to.")
)

func main() {
	controller := controller.Controller{}
	controller.Connect(*addr)
	for {
		cmd, err := controller.ReadCommandRequest()
		if err != nil {
			log.Printf("Failed to read command: %v", err)
		}
		switch cmd.Type {
		case pb.ECHO_CMD_TYPE:
			resp := modules.RunEchoCommand(cmd.GetEchoCommandRequest())
			controller.WriteCommandResponse(resp)
		}
	}
}
