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

type Config struct {
	AgentServerPort        int     `json:"agentServerPort"`
	AgentManagerServerPort int     `json:"agentManagerServerPort"`
	KeyPair                KeyPair `json:"keyPair"`
	AgentCert              string  `json:"agentCert"`
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
