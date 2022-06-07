package modules

import (
	"context"
	"log"
	"strings"
	pb "viper/protos/cmds"
)

var (
	requests  = make(chan string)
	responses = make(chan []byte)
)

func (s *AgentServer) RunShellCommand(stream pb.Agent_RunShellCommandServer) error {
	for {
		cmd := <-requests
		log.Printf("Sending command to the agent: '%s'", cmd)
		stream.Send(&pb.ShellCommandRequest{Cmd: strings.TrimSpace(cmd)})
		in, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive shell command output: %v", err)
		}
		responses <- in.Output
	}
}

func (s *AgentManagerServer) RunShellCommand(context context.Context, req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	log.Printf("Sending command to the agent server: '%s'", req.Cmd)
	requests <- req.Cmd
	response := pb.ShellCommandResponse{Output: <-responses}
	log.Printf("Received command response.")
	return &response, nil
}
