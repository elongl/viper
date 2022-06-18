package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"

	_ "embed"
	"viper/controller/agents"
	"viper/controller/commands"
	pb "viper/protos/cmds"

	"google.golang.org/grpc"
)

var (
	agentServerPort        = flag.Int("port", 50051, "Agent server port")
	agentManagerServerPort = flag.Int("management-port", 50052, "Agent management server port")

	//go:embed certs/controller.cert
	certBuffer []byte
	//go:embed certs/controller.key
	keyBuffer []byte
	//go:embed certs/agent.cert
	agentCertBuffer []byte
)

func runAgentServer() {
	cert, err := tls.X509KeyPair(certBuffer, keyBuffer)
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(agentCertBuffer)
	if !ok {
		log.Fatalf("Failed to parse agent's client certificate.")
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: caCertPool}
	lis, err := tls.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *agentServerPort), tlsCfg)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Agent server listening at %v", lis.Addr())
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
		}
		go agents.InitAgent(conn)
	}
}

func runAgentManagerServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *agentManagerServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentManagerServer(server, &commands.AgentManagerServer{})
	log.Printf("Agent manager server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func main() {
	flag.Parse()
	go runAgentServer()
	go runAgentManagerServer()
	select {}
}
