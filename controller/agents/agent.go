package agents

import (
	"errors"
	"log"
	"net"
)

type Agent struct {
	Hostname string
	Conn     net.Conn
	Alive    bool
}

var (
	agents []Agent
)

func GetAgent(hostname string) (*Agent, error) {
	for _, agent := range agents {
		if agent.Hostname == hostname {
			return &agent, nil
		}
	}
	return nil, errors.New("Agent not found")
}

func InitAgent(conn net.Conn) {
	log.Printf("Initializing agent: %v", conn.RemoteAddr())
	agent := Agent{Conn: conn, Hostname: "egk", Alive: true}
	agents = append(agents, agent)
}

func DeleteAgent(agent *Agent) {
	log.Printf("Deleting agent: %v", agent)
	agent.Alive = false
}
