package agents

import (
	"errors"
	"log"
	"net"
)

type ShellCommandQueue struct {
	Reqs  chan string
	Resps chan []byte
}

type EchoCommandQueue struct {
	Reqs  chan string
	Resps chan string
}

type AgentQueues struct {
	Shell ShellCommandQueue
	Echo  EchoCommandQueue
}

type Agent struct {
	Hostname string
	Conn     net.Conn
	Queues   AgentQueues
	Alive    bool
}

var (
	Agents []Agent
)

func initQueue(agent *Agent) {
	agent.Queues.Shell.Reqs = make(chan string)
	agent.Queues.Shell.Resps = make(chan []byte)

	agent.Queues.Echo.Reqs = make(chan string)
	agent.Queues.Echo.Resps = make(chan string)
}

func GetAgent(hostname string) (*Agent, error) {
	for _, agent := range Agents {
		if agent.Hostname == hostname {
			return &agent, nil
		}
	}
	return nil, errors.New("Agent not found")
}

func InitAgent(conn net.Conn) {
	log.Printf("Initializing agent: %v", conn.RemoteAddr())
	agent := Agent{Conn: conn}
	initQueue(&agent)
	Agents = append(Agents, agent)
}

func DeleteAgent(agent *Agent) {
	log.Printf("Deleting agent: %v", agent)
	close(agent.Queues.Shell.Reqs)
	close(agent.Queues.Shell.Resps)
	agent.Alive = false
}
