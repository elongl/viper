package commands

import (
	"context"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) RunShellCommand(ctx context.Context, req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	log.Printf("[%d] Sending shell command.", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.RunShellCommand(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Received shell response.")
	return resp, nil
}
