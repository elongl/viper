package commands

import (
	"context"
	"log"
	"viper/controller/agents"
	"viper/protos/cmds"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) RunEchoCommand(ctx context.Context, req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	log.Printf("Sending echo command to agent: '%s'", req.Hostname)
	agent, err := agents.GetAgent(req.Hostname)
	if err != nil {
		return nil, err
	}
	cmdReq := &pb.CommandRequest{Type: cmds.ECHO_CMD_TYPE, Req: &pb.CommandRequest_EchoCommandRequest{EchoCommandRequest: req}}
	err = agent.Write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.EchoCommandResponse{}
	err = agent.Read(resp)
	if err != nil {
		return nil, err
	}
	log.Printf("Received echo response.")
	return resp, nil
}
