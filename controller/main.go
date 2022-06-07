package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"viper/controller/modules"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	agentServerPort        = flag.Int("port", 50051, "Agent server port")
	agentManagerServerPort = flag.Int("management-port", 50052, "Agent management server port")
)

func runAgentServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *agentServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentServer(server, &modules.AgentServer{})
	log.Printf("Agent server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func runAgentManagerServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *agentManagerServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentManagerServer(server, &modules.AgentManagerServer{})
	log.Printf("Agent manager server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func main() {
	flag.Parse()
	go runAgentServer()
	go runAgentManagerServer()
	select {}
}
