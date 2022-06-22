package modules

import (
	"log"
	"net"
	pb "viper/protos/cmds"

	socks "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
)

var (
	streamCh = make(chan net.Conn)
	stopCh   = make(chan bool)
)

func StartSocksServer(req *pb.StartSocksServerRequest, controllerSession *yamux.Session) *pb.StartSocksServerResponse {
	log.Printf("Starting SOCKS server.")
	// TODO: This goroutine is currently leaking (never exits). Fix.
	go acceptConnsToChans(controllerSession)
	go serveConns()
	return &pb.StartSocksServerResponse{}
}

func StopSocksServer(req *pb.StopSocksServerRequest, controllerSession *yamux.Session) *pb.StopSocksServerResponse {
	log.Printf("Stopping SOCKS server.")
	stopCh <- true
	return &pb.StopSocksServerResponse{}
}

func acceptConnsToChans(controllerSession *yamux.Session) {
	for {
		stream, err := controllerSession.Accept()
		if err != nil {
			log.Printf("Failed to accept SOCKS connection: %v", err)
			return
		}
		log.Printf("Received controller's SOCKS connection.")
		streamCh <- stream
	}
}

func serveConns() {
	for {
		select {
		case <-stopCh:
			log.Print("Stopping the SOCKS server.")
			return
		case stream := <-streamCh:
			server, err := socks.New(&socks.Config{})
			if err != nil {
				log.Printf("Failed to create a SOCKS server: %v", err)
				return
			}
			go server.ServeConn(stream)
		}
	}
}
