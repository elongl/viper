package modules

import (
	"log"
	pb "viper/protos/cmds"
)

func RunEchoCommand(req *pb.EchoCommandRequest) *pb.EchoCommandResponse {
	log.Printf("Running echo command: '%s'.", req.Data)
	return &pb.EchoCommandResponse{Data: req.Data}
}
