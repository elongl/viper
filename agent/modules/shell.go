//go:build !windows

package modules

import (
	"log"
	"os/exec"
	pb "viper/protos/cmds"
)

func RunShellCommand(req *pb.ShellCommandRequest) *pb.ShellCommandResponse {
	log.Printf("running shell command: '%s'", req.Cmd)
	cmd := exec.Command("sh", "-c", req.Cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("failed to run shell command: %v", err)
		return &pb.ShellCommandResponse{Data: out, Err: err.Error()}
	}
	return &pb.ShellCommandResponse{Data: out}
}
