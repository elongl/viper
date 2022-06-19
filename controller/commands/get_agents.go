package commands

import (
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) GetAgents(req *pb.Empty, stream pb.AgentManager_GetAgentsServer) error {
	log.Printf("Getting the agents.")
	for _, agent := range agents.Agents {
		err := stream.Send(&pb.AgentInfo{Id: agent.Id, Alive: agent.Alive, ConnectTime: agent.ConnectTime.String()})
		if err != nil {
			return err
		}
	}
	log.Printf("Sent the agents.")
	return nil
}
