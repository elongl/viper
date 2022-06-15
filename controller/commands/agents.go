package commands

import (
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) GetAgents(in *pb.Empty, stream pb.AgentManager_GetAgentsServer) error {
	log.Printf("Getting the available agents.")
	for agentAddr := range agents.Agents {
		stream.Send(&pb.AgentInfo{Addr: agentAddr})
	}
	return nil
}
