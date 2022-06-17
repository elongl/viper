package modules

import (
	"log"
	pb "viper/protos/cmds"
)

func RunEchoCommand(req *pb.EchoCommandRequest) *pb.EchoCommandResponse {
	log.Print("Running echo command.")
	return &pb.EchoCommandResponse{Text: req.Text}
}
