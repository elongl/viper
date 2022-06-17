package controller

import (
	"encoding/binary"
	"errors"
	"fmt"
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
		msg := fmt.Sprintf("Failed to read command size: %v", err)
		return nil, errors.New(msg)
	}

	cmdBuffer := make([]byte, cmdSize)
	_, err = cnc.conn.Read(cmdBuffer)
	if err != nil {
		msg := fmt.Sprintf("Failed to read command: %v", err)
		return nil, errors.New(msg)
	}
	cmd := &pb.CommandRequest{}
	err = proto.Unmarshal(cmdBuffer, cmd)
	if err != nil {
		msg := fmt.Sprintf("Failed to unmarshal command: %v", err)
		return nil, errors.New(msg)
	}
	return cmd, nil
}

func (cnc *Controller) WriteCommandResponse(resp proto.Message) error {
	respBuffer, err := proto.Marshal(resp)
	if err != nil {
		msg := fmt.Sprintf("Failed to marshal response: %v", err)
		return errors.New(msg)
	}
	respBufferLen := int64(len(respBuffer))
	err = binary.Write(cnc.conn, binary.LittleEndian, &respBufferLen)
	if err != nil {
		msg := fmt.Sprintf("Failed to write response size: %v", err)
		return errors.New(msg)
	}
	_, err = cnc.conn.Write(respBuffer)
	if err != nil {
		msg := fmt.Sprintf("Failed to write response: %v", err)
		return errors.New(msg)
	}
	return nil
}
