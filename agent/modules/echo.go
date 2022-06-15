package modules

import (
	"context"
	"log"
	pb "viper/protos/cmds"
)

func runEchoModule(client pb.AgentClient) {
	stream, err := client.RunEchoCommand(context.Background())
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}
	for {
		in, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive an echo command: %v", err)
		}
		log.Printf("Received echo command: '%s'", in.Text)
		err = stream.Send(&pb.EchoCommandResponse{Text: in.Text})
		if err != nil {
			log.Printf("Failed to send echo command response: %v", err)
		}
	}
}
