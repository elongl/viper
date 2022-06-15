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
	resp, err := client.RunShellCommand(ctx, &pb.ShellCommandRequest{Cmd: "whoami", Addr: "127.0.0.1:59851"})
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	log.Printf("Received command response: '%s'", strings.TrimSpace(string(resp.Output)))
}
