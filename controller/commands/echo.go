package commands

import (
	"context"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) RunEchoCommand(ctx context.Context, req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	log.Printf("Sending echo command to agent: '%s'", req.Hostname)
	agent, err := agents.GetAgent(req.Hostname)
	if err != nil {
		return nil, err
	}
	resp, err := agent.RunEchoCommand(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Received echo response.")
	return resp, nil
}
