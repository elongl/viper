package commands

import (
	pb "viper/protos/cmds"
)

type AgentManagerServer struct {
	pb.UnimplementedAgentManagerServer
}
