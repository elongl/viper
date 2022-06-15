package agents

import (
	"errors"
	"fmt"
	"log"
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
	Addr   string
	Queues AgentQueues
}

var (
	Agents = make(map[string]Agent)
)

func initQueue(agent *Agent) {
	agent.Queues.Shell.Reqs = make(chan string)
	agent.Queues.Shell.Resps = make(chan []byte)

	agent.Queues.Echo.Reqs = make(chan string)
	agent.Queues.Echo.Resps = make(chan string)
}

func GetAgent(peerAddr string) (*Agent, error) {
	if agent, ok := Agents[peerAddr]; ok {
		return &agent, nil
	} else {
		msg := fmt.Sprintf("Agent '%s' is not connected", peerAddr)
		return nil, errors.New(msg)
	}
}

func InitAgent(peerAddr string) {
	if _, ok := Agents[peerAddr]; !ok {
		log.Printf("Initializing agent: %v", peerAddr)
		agent := Agent{Addr: peerAddr}
		initQueue(&agent)
		Agents[peerAddr] = agent
	}
}

func DeleteAgent(peerAddr string) {
	log.Printf("Deleting agent: %v", peerAddr)
	agent := Agents[peerAddr]
	close(agent.Queues.Shell.Reqs)
	close(agent.Queues.Shell.Resps)
	delete(Agents, peerAddr)
}
