package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	shellCmds = make(chan string)
)

type agentServer struct {
	pb.UnimplementedAgentServer
}

func (s *agentServer) RunShellCommand(stream pb.Agent_RunShellCommandServer) error {
	shellCmds <- "whoami"
	for {
		stream.Send(&pb.ShellCommandRequest{Cmd: strings.TrimSpace(<-shellCmds)})
		in, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive shell command output: %v", err)
		}
		fmt.Printf(">> %s", in.Output)
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
