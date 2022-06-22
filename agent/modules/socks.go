package modules

import (
	"log"
	"net"
	pb "viper/protos/cmds"

	socks "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
)

func StartSocksServer(req *pb.StartSocksServerRequest, controllerConn net.Conn) *pb.StartSocksServerResponse {
	log.Printf("Starting SOCKS server.")
	session, err := yamux.Server(controllerConn, nil)
	if err != nil {
		return &pb.StartSocksServerResponse{Err: err.Error()}
	}
	go func() {
		for {
			stream, err := session.Accept()
			if err != nil {
				log.Printf("Failed to accept connection: %v", err)
				return
			}
			log.Print("Received controller's connection.")
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
