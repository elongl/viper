package modules

import (
	"log"
	pb "viper/protos/cmds"
)

func RunEchoCommand(req *pb.EchoCommandRequest) *pb.EchoCommandResponse {
	log.Printf("running echo command: '%s'", req.Data)
	return &pb.EchoCommandResponse{Data: req.Data}
}
