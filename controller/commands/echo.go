package commands

import (
	"context"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"

	"google.golang.org/grpc/peer"
)

func (s *AgentServer) RunEchoCommand(stream pb.Agent_RunEchoCommandServer) error {
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		log.Fatal("Failed to get peer from context")
	}
	peerAddr := peer.Addr.String()
	queue := agents.Agents[peerAddr].Queues.Echo
	for {
		text := <-queue.Reqs
		log.Printf("Sending echo to the agent: '%s'", text)
		stream.Send(&pb.EchoCommandRequest{Text: text})
		in, err := stream.Recv()
		if err != nil {
			agents.DeleteAgent(peerAddr)
			return err
		}
		queue.Resps <- in.Text
	}
}

func (s *AgentManagerServer) RunEchoCommand(ctx context.Context, req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	log.Printf("Sending echo command to the agent server: '%s'", req.Text)
	agent, err := agents.GetAgent(req.Addr)
	if err != nil {
		return nil, err
	}
	queue := agent.Queues.Echo
	queue.Reqs <- req.Text
	response := pb.EchoCommandResponse{Text: <-queue.Resps}
	log.Printf("Received echo response.")
	return &response, nil
}
