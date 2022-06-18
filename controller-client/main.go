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

	echoOutput, err := client.RunEchoCommand(ctx, &pb.EchoCommandRequest{Text: "Hello!", AgentId: 0})
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	log.Printf("Received echo response: '%s'", echoOutput.Text)

	shellOutput, err := client.RunShellCommand(ctx, &pb.ShellCommandRequest{Cmd: "whoami", AgentId: 0})
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	log.Printf("Received shell response: '%s'", strings.TrimSpace(string(shellOutput.Output)))
}
