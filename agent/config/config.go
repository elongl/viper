package config

import (
	_ "embed"
	"encoding/json"
	"log"
)

type KeyPair struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type PersistenceConfig struct {
	Path     string `json:"path"`
	TaskName string `json:"taskName"`
}

type Config struct {
	ControllerAddress string            `json:"controllerAddress"`
	Persistence       PersistenceConfig `json:"persistence"`
	KeyPair           KeyPair           `json:"keyPair"`
}

var (
	//go:embed config.json
	configBuffer []byte
	Conf         = getConfig()
)

func getConfig() *Config {
	cfg := &Config{}
	err := json.Unmarshal(configBuffer, cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	return cfg
}
