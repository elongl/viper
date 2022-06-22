package commands

import (
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) GetAgents(req *pb.GetAgentsRequest, stream pb.AgentManager_GetAgentsServer) error {
	log.Printf("getting the agents")
	for _, agent := range agents.Agents {
		agentAlive := agent.IsAlive()
		if req.AliveOnly && !agentAlive {
			continue
		}
		err := stream.Send(&pb.AgentInfo{Id: agent.Id, Alive: agentAlive, ConnectTime: agent.ConnectTime.String()})
		if err != nil {
			return err
		}
	}
	log.Printf("sent the agents")
	return nil
}
