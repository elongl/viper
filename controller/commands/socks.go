package commands

import (
	"context"
	"controller/agents"
	pb "controller/protos/cmds"
	"log"
)

func (s *AgentManagerServer) StartSocksServer(ctx context.Context, req *pb.StartSocksServerRequest) (*pb.StartSocksServerResponse, error) {
	log.Printf("[%d] starting SOCKS server", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.StartSocksServer(req)
	if err != nil {
		return nil, err
	}
	log.Printf("started SOCKS server")
	return resp, nil
}

func (s *AgentManagerServer) StopSocksServer(ctx context.Context, req *pb.StopSocksServerRequest) (*pb.StopSocksServerResponse, error) {
	log.Printf("[%d] stopping SOCKS server", req.AgentId)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.StopSocksServer(req)
	if err != nil {
		return nil, err
	}
	log.Printf("stopped SOCKS server")
	return resp, nil
}
