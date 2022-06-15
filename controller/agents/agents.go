package agents

import "log"

type ShellCommandQueue struct {
	Reqs  chan string
	Resps chan []byte
}

type AgentQueues struct {
	Shell ShellCommandQueue
}

type Agent struct {
	Queues AgentQueues
}

var (
	Agents = make(map[string]Agent)
)

func initShellQueue(agent *Agent) {
	agent.Queues.Shell.Reqs = make(chan string)
	agent.Queues.Shell.Resps = make(chan []byte)
}

func InitAgent(peerAddr string) {
	log.Printf("Initializing agent: %v", peerAddr)
	agent := Agent{}
	initShellQueue(&agent)
	Agents[peerAddr] = agent
}

func DeleteAgent(peerAddr string) {
	log.Printf("Deleting agent: %v", peerAddr)
	delete(Agents, peerAddr)
}
