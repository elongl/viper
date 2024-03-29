package commands

import (
	"context"
	"controller/agents"
	pb "controller/protos/cmds"
	"log"
)

func (s *AgentManagerServer) RunEchoCommand(ctx context.Context, req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	log.Printf("[%d] sending echo command", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.RunEchoCommand(req)
	if err != nil {
		return nil, err
	}
	log.Printf("received echo response")
	return resp, nil
}
