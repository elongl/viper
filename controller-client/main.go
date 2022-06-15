package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	pb "viper/protos/cmds"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "The address to connect to.")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAgentManagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	agentsStream, err := client.GetAgents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to get agents: %v", err)
	}
	for {
		agent, err := agentsStream.Recv()
		if err != nil {
			break
		}
		log.Printf("Agent: %s", agent.GetAddr())
	}

	agentAddr := "127.0.0.1:61332"

	shellOutput, err := client.RunShellCommand(ctx, &pb.ShellCommandRequest{Cmd: "whoami", Addr: agentAddr})
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	log.Printf("Received command response: '%s'", strings.TrimSpace(string(shellOutput.Output)))

	echoOutput, err := client.RunEchoCommand(ctx, &pb.EchoCommandRequest{Text: "Hello World!", Addr: agentAddr})
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	log.Printf("Received echo response: '%s'", echoOutput.Text)
}
