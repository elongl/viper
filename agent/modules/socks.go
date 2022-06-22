package modules

import (
	"log"
	pb "viper/protos/cmds"

	socks "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
)

func StartSocksServer(req *pb.StartSocksServerRequest, controllerSession *yamux.Session) *pb.StartSocksServerResponse {
	log.Printf("Starting SOCKS server.")
	go func() {
		for {
			stream, err := controllerSession.Accept()
			if err != nil {
				log.Printf("Failed to accept SOCKS connection: %v", err)
				return
			}
			log.Print("Received controller's SOCKS connection.")
			server, err := socks.New(&socks.Config{})
			if err != nil {
				log.Printf("Failed to create a SOCKS server: %v", err)
				return
			}
			go server.ServeConn(stream)
		}
	}()
	return &pb.StartSocksServerResponse{}
}
