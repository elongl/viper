package controller

import (
	"encoding/binary"
	"log"
	"net"
	"time"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

type Controller struct {
	conn net.Conn
}

func (cnc *Controller) Connect(addr string) {
	for {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Printf("Failed to connect to controller: %v", err)
			time.Sleep(time.Minute * 5)
		} else {
			log.Print("Connected to controller.")
			cnc.conn = conn
			return
		}
	}
}

func (cnc *Controller) ReadCommandRequest() (*pb.CommandRequest, error) {
	var cmdSize int64
	err := binary.Read(cnc.conn, binary.LittleEndian, &cmdSize)
	if err != nil {
		return nil, err
	}

	cmdBuffer := make([]byte, cmdSize)
	_, err = cnc.conn.Read(cmdBuffer)
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

func (cnc *Controller) WriteCommandResponse(resp proto.Message) error {
	respBuffer, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	respBufferLen := int64(len(respBuffer))
	err = binary.Write(cnc.conn, binary.LittleEndian, &respBufferLen)
	if err != nil {
		return err
	}
	_, err = cnc.conn.Write(respBuffer)
	if err != nil {
		return err
	}
	return nil
}
