package controller

import (
	"crypto/tls"
	_ "embed"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
	"viper"
	pb "viper/protos/cmds"

	"github.com/hashicorp/yamux"
	"google.golang.org/protobuf/proto"
)

type Controller struct {
	Addr      string
	Session   *yamux.Session
	cmdStream net.Conn
}

func (cnc *Controller) Connect() {
	certBuffers := viper.Conf.Agent.Cert
	cert, err := tls.X509KeyPair([]byte(certBuffers.Cert), []byte(certBuffers.Key))
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	for {
		log.Print("Connecting to controller.")
		conn, err := tls.Dial("tcp", cnc.Addr, tlsCfg)
		if err != nil {
			log.Printf("Failed to connect to controller: %v", err)
			time.Sleep(time.Minute * 1)
			continue
		}
		log.Print("Connected to controller.")
		session, err := yamux.Client(conn, nil)
		if err != nil {
			log.Printf("Failed to create multiplexing client: %v", err)
			continue
		}
		cnc.Session = session
		stream, err := session.Open()
		if err != nil {
			log.Printf("Failed to open a multiplexed stream:", err)
			continue
		}
		cnc.cmdStream = stream
		return
	}
}

func (cnc *Controller) ReadCommandRequest() (*pb.CommandRequest, error) {
	for {
		var cmdSize int64
		err := binary.Read(cnc.cmdStream, binary.LittleEndian, &cmdSize)
		if err != nil {
			cnc.Connect()
			continue
		}
		cmdBuffer := make([]byte, cmdSize)
		_, err = io.ReadFull(cnc.cmdStream, cmdBuffer)
		if err != nil {
			cnc.Connect()
			continue
		}
		cmd := &pb.CommandRequest{}
		err = proto.Unmarshal(cmdBuffer, cmd)
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
}

func (cnc *Controller) WriteCommandResponse(resp proto.Message) error {
	for {
		respBuffer, err := proto.Marshal(resp)
		if err != nil {
			return err
		}
		respBufferLen := int64(len(respBuffer))
		err = binary.Write(cnc.cmdStream, binary.LittleEndian, &respBufferLen)
		if err != nil {
			cnc.Connect()
			continue
		}
		_, err = cnc.cmdStream.Write(respBuffer)
		if err != nil {
			cnc.Connect()
			continue
		}
		return nil
	}
}
