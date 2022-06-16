package commands

import (
	"context"
	"encoding/binary"
	"log"
	"viper/controller/agents"
	"viper/protos/cmds"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

func (s *AgentManagerServer) RunEchoCommand(ctx context.Context, req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	log.Printf("Sending echo command to agent: '%s'", req.Hostname)
	agent, err := agents.GetAgent(req.Hostname)
	if err != nil {
		return nil, err
	}
	cmd := &pb.Command{Type: cmds.ECHO_CMD_TYPE, Command: &pb.Command_EchoCommandRequest{EchoCommandRequest: req}}
	cmdBuffer, err := proto.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	cmdBufferLen := int64(len(cmdBuffer))
	err = binary.Write(agent.Conn, binary.LittleEndian, &cmdBufferLen)
	if err != nil {
		return nil, err
	}
	_, err = agent.Conn.Write(cmdBuffer)
	if err != nil {
		return nil, err
	}
	var respSize int64
	err = binary.Read(agent.Conn, binary.LittleEndian, &respSize)
	if err != nil {
		return nil, err
	}
	respBuffer := make([]byte, respSize)
	_, err = agent.Conn.Read(respBuffer)
	if err != nil {
		return nil, err
	}
	resp := &pb.EchoCommandResponse{}
	err = proto.Unmarshal(respBuffer, resp)
	if err != nil {
		return nil, err
	}
	log.Printf("Received echo response.")
	return resp, nil
}
