package viper

import (
	_ "embed"
	"encoding/json"
	"log"
)

type CertificateConfig struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type AgentConfig struct {
	ControllerAddress string            `json:"controllerAddress"`
	PersistencePath   string            `json:"persistencePath"`
	Cert              CertificateConfig `json:"cert"`
}

type ControllerConfig struct {
	AgentServerPort        int               `json:"agentServerPort"`
	AgentManagerServerPort int               `json:"agentManagerServerPort"`
	Cert                   CertificateConfig `json:"cert"`
}

type Config struct {
	Agent      AgentConfig      `json:"agent"`
	Controller ControllerConfig `json:"controller"`
}

var (
	//go:embed config.json
	configBuffer []byte

	Conf = getConfig()
)

func getConfig() *Config {
	cfg := &Config{}
	err := json.Unmarshal(configBuffer, cfg)
	if err != nil {
		log.Fatal("Failed to parse config: ", err)
	}
	return cfg
}
