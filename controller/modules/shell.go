package modules

import (
	"context"
	"log"
	"strings"
	pb "viper/protos/cmds"

	"google.golang.org/grpc/peer"
)

var (
	requests  = make(map[string](chan string))
	responses = make(map[string](chan []byte))
)

func (s *AgentServer) RunShellCommand(stream pb.Agent_RunShellCommandServer) error {
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		log.Fatal("Failed to get peer from context")
	}
	peerAddr := peer.Addr.String()
	requests[peerAddr] = make(chan string)
	responses[peerAddr] = make(chan []byte)
	for {
		log.Printf("Listening on :%v", peerAddr)
		cmd := <-requests[peerAddr]
		log.Printf("Sending command to the agent: '%s'", cmd)
		stream.Send(&pb.ShellCommandRequest{Cmd: strings.TrimSpace(cmd)})
		in, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive shell command output: %v", err)
		}
		responses[peerAddr] <- in.Output
	}
}

func (s *AgentManagerServer) RunShellCommand(context context.Context, req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	log.Printf("Sending command to the agent server: '%s'", req.Cmd)
	requests[req.Addr] <- req.Cmd
	response := pb.ShellCommandResponse{Output: <-responses[req.Addr]}
	log.Printf("Received command response.")
	return &response, nil
}
