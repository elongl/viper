package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"viper/controller/agents"
	"viper/controller/commands"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	agentServerPort        = flag.Int("port", 50051, "Agent server port")
	agentManagerServerPort = flag.Int("management-port", 50052, "Agent management server port")
)

func runAgentServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *agentServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Agent server listening at %v", lis.Addr())
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
		}
		go agents.InitAgent(conn)
	}
}

func runAgentManagerServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *agentManagerServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentManagerServer(server, &commands.AgentManagerServer{})
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
