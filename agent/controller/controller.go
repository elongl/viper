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

	"google.golang.org/protobuf/proto"
)

const (
	ERR_CONTROLLER_CONN_CLOSED = "Connection to controller has been closed."
)

type Controller struct {
	Addr string
	conn net.Conn
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
		} else {
			log.Print("Connected to controller.")
			cnc.conn = conn
			return
		}
	}
}

func (cnc *Controller) ReadCommandRequest() (*pb.CommandRequest, error) {
	for {
		var cmdSize int64
		err := binary.Read(cnc.conn, binary.LittleEndian, &cmdSize)
		if err == io.EOF {
			log.Fatal(ERR_CONTROLLER_CONN_CLOSED)
			continue
		}
		if err != nil {
			return nil, err
		}

		cmdBuffer := make([]byte, cmdSize)
		_, err = cnc.conn.Read(cmdBuffer)
		if err == io.EOF {
			log.Fatal(ERR_CONTROLLER_CONN_CLOSED)
			continue
		}
		if err != nil {
			return nil, err
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
		err = binary.Write(cnc.conn, binary.LittleEndian, &respBufferLen)
		if err == io.EOF {
			log.Fatal("Controller has closed the connection.")
			continue
		}
		if err != nil {
			return err
		}
		_, err = cnc.conn.Write(respBuffer)
		if err == io.EOF {
			log.Fatal(ERR_CONTROLLER_CONN_CLOSED)
			continue
		}
		if err != nil {
			return err
		}
		return nil
	}
}
