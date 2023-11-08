package commands

import (
	pb "controller/protos/cmds"
)

type AgentManagerServer struct {
	pb.UnimplementedAgentManagerServer
}
