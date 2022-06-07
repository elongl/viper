package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	agentServerPort        = flag.Int("port", 50051, "Agent server port")
	agentManagerServerPort = flag.Int("port", 50052, "Agent management server port")

	requests  = make(chan string)
	responses = make(chan string)
)

type agentServer struct {
	pb.UnimplementedAgentServer
}

type agentManagerServer struct {
	pb.UnimplementedAgentManagerServer
}

func (s *agentServer) RunShellCommand(stream pb.Agent_RunShellCommandServer) error {
	for {
		cmd := <-requests
		log.Printf("Sending command to the agent: '%s'", cmd)
		stream.Send(&pb.ShellCommandRequest{Cmd: strings.TrimSpace(cmd)})
		in, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive shell command output: %v", err)
		}
		responses <- in.Output
	}
}

func (s *agentManagerServer) RunShellCommand(context context.Context, req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	log.Printf("Sending command to the agent server: '%s'", req.Cmd)
	requests <- req.Cmd
	response := pb.ShellCommandResponse{Output: <-responses}
	log.Printf("Received command response: '%s'", response.Output)
	return &response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *agentServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentServer(server, &agentServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
