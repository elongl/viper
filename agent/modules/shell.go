package modules

import (
	"log"
	"os/exec"
	"runtime"
	pb "viper/protos/cmds"
)

func RunShellCommand(req *pb.ShellCommandRequest) *pb.ShellCommandResponse {
	var cmd *exec.Cmd
	log.Printf("Running shell command: '%s'.", req.Cmd)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", req.Cmd)
		// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("sh", "-c", req.Cmd)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run shell command: %v", err)
		return &pb.ShellCommandResponse{Err: err.Error()}
	}
	return &pb.ShellCommandResponse{Output: out}
}
