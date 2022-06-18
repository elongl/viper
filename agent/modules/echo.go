package modules

import (
	"log"
	pb "viper/protos/cmds"
)

func RunEchoCommand(req *pb.EchoCommandRequest) *pb.EchoCommandResponse {
	log.Printf("Running echo command: '%s'.", req.Text)
	return &pb.EchoCommandResponse{Text: req.Text}
}
