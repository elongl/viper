package main

import (
	"flag"
	"log"

	"viper/agent/modules/shell"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "The address to connect to.")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAgentClient(conn)
	go shell.RunModule(client)
	log.Print("Agent started.")
	select {}
}
