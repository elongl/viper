package commands

import (
	"context"
	"fmt"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) Screenshot(ctx context.Context, req *pb.ScreenshotRequest) (*pb.ScreenshotResponse, error) {
	log.Printf("[%d] Sending screenshot command.", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.Screenshot(req)
	if err != nil {
		return nil, err
	}
	if resp.Err != "" {
		return nil, fmt.Errorf("Failed to take screenshot: %v", resp.Err)
	}
	log.Printf("Received screenshot response.")
	return resp, nil
}
