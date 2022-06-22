package modules

import (
	"log"
	"net"
	pb "viper/protos/cmds"

	socks "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
)

var (
	streamCh         = make(chan net.Conn)
	stopCh           = make(chan bool)
	startedAccepting = false
)

func StartSocksServer(req *pb.StartSocksServerRequest, controllerSession *yamux.Session) *pb.StartSocksServerResponse {
	log.Printf("starting SOCKS server")
	if !startedAccepting {
		go acceptConnsToChans(controllerSession)
		startedAccepting = true
	}
	go serveConns()
	return &pb.StartSocksServerResponse{}
}

func StopSocksServer(req *pb.StopSocksServerRequest, controllerSession *yamux.Session) *pb.StopSocksServerResponse {
	log.Printf("stopping SOCKS server")
	stopCh <- true
	return &pb.StopSocksServerResponse{}
}

func acceptConnsToChans(controllerSession *yamux.Session) {
	for {
		stream, err := controllerSession.Accept()
		if err != nil {
			log.Printf("failed to accept SOCKS connection: %v", err)
			return
		}
		log.Printf("received controller's SOCKS connection")
		streamCh <- stream
	}
}

func serveConns() {
	for {
		select {
		case <-stopCh:
			log.Print("stopping the SOCKS server")
			return
		case stream := <-streamCh:
			server, err := socks.New(&socks.Config{})
			if err != nil {
				log.Printf("failed to create a SOCKS server: %v", err)
				return
			}
			go server.ServeConn(stream)
		}
	}
}
