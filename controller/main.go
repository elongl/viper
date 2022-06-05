package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type agentServer struct {
	pb.UnimplementedAgentServer
}

func (s *agentServer) RunShellCommand(stream pb.Agent_RunShellCommandServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received shell command output: '%s'", in.Output)
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
