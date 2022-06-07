package modules

import (
	pb "viper/protos/cmds"
)

type AgentServer struct {
	pb.UnimplementedAgentServer
}

type AgentManagerServer struct {
	pb.UnimplementedAgentManagerServer
}
