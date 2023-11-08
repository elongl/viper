package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"

	"controller/agents"
	"controller/commands"
	"controller/config"
	pb "controller/protos/cmds"

	"google.golang.org/grpc"
)

const (
	MAX_MSG_LEN = 100 * 1024 * 1024
)

func runAgentServer() {
	keyPairBuffers := config.Conf.KeyPair
	keyPair, err := tls.X509KeyPair([]byte(keyPairBuffers.Cert), []byte(keyPairBuffers.Key))
	if err != nil {
		log.Fatalf("failed to load certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM([]byte(config.Conf.AgentCert))
	if !ok {
		log.Fatalf("failed to parse agent's client certificate")
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{keyPair}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: caCertPool}
	lis, err := tls.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Conf.AgentServerPort), tlsCfg)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("agent server listening at %v", lis.Addr())
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("failed to accept connection: %v", err)
		}
		agents.InitAgent(conn)
	}
}

func runAgentManagerServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", config.Conf.AgentManagerServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(grpc.MaxRecvMsgSize(MAX_MSG_LEN), grpc.MaxSendMsgSize(MAX_MSG_LEN))
	pb.RegisterAgentManagerServer(server, &commands.AgentManagerServer{})
	log.Printf("agent manager server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	flag.Parse()
	go runAgentServer()
	go runAgentManagerServer()
	select {}
}
