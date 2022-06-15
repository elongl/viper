package commands

import (
	"context"
	"log"
	"strings"
	"viper/controller/agents"
	pb "viper/protos/cmds"

	"google.golang.org/grpc/peer"
)

func (s *AgentServer) RunShellCommand(stream pb.Agent_RunShellCommandServer) error {
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		log.Fatal("Failed to get peer from context")
	}
	peerAddr := peer.Addr.String()
	queue := agents.Agents[peerAddr].Queues.Shell
	for {
		cmd := <-queue.Reqs
		log.Printf("Sending command to the agent: '%s'", cmd)
		stream.Send(&pb.ShellCommandRequest{Cmd: strings.TrimSpace(cmd)})
		in, err := stream.Recv()
		if err != nil {
			agents.DeleteAgent(peerAddr)
			return err
		}
		queue.Resps <- in.Output
	}
}

func (s *AgentManagerServer) RunShellCommand(ctx context.Context, req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	log.Printf("Sending command to the agent server: '%s'", req.Cmd)
	agent, err := agents.GetAgent(req.Addr)
	if err != nil {
		return nil, err
	}
	queue := agent.Queues.Shell
	queue.Reqs <- req.Cmd
	response := pb.ShellCommandResponse{Output: <-queue.Resps}
	log.Printf("Received command response.")
	return &response, nil
}
