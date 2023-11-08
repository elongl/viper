//go:build !windows

package modules

import (
	pb "agent/protos/cmds"
	"fmt"
	"log"
	"os/exec"
)

func RunShellCommand(req *pb.ShellCommandRequest) *pb.ShellCommandResponse {
	log.Printf("running shell command: '%s'", req.Cmd)
	cmd := exec.Command("sh", "-c", req.Cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := fmt.Sprintf("%v : %s", err, out)
		log.Printf("failed to run shell command: %s", msg)
		return &pb.ShellCommandResponse{Data: out, Err: msg}
	}
	return &pb.ShellCommandResponse{Data: out}
}
