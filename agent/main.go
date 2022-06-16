package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net"
	"viper/protos/cmds"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "The address to connect to.")
)

func main() {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	log.Print("Connected to controller.")
	for {
		var cmdSize int64
		err := binary.Read(conn, binary.LittleEndian, &cmdSize)
		if err != nil {
			log.Fatalf("Failed to read: %v", err)
		}

		cmdBuffer := make([]byte, cmdSize)
		_, err = conn.Read(cmdBuffer)
		if err != nil {
			log.Fatalf("Failed to read: %v", err)
		}
		cmd := &pb.Command{}
		err = proto.Unmarshal(cmdBuffer, cmd)
		if err != nil {
			log.Fatalf("Failed to unmarshal command: %v", err)
		}

		switch cmd.Type {
		case cmds.ECHO_CMD_TYPE:
			log.Print("Received echo command.")
			echoCmd := cmd.GetEchoCommandRequest()
			log.Printf("Echo's text: '%s'", echoCmd.Text)
			resp := &pb.EchoCommandResponse{Text: echoCmd.Text}
			respBuffer, err := proto.Marshal(resp)
			if err != nil {
				log.Fatalf("Failed to marshal response: %v", err)
			}
			respBufferLen := int64(len(respBuffer))
			err = binary.Write(conn, binary.LittleEndian, &respBufferLen)
			if err != nil {
				log.Fatalf("Failed to write: %v", err)
			}
			_, err = conn.Write(respBuffer)
			if err != nil {
				log.Fatalf("Failed to write: %v", err)
			}
			log.Print("Sent echo response.")
		}
	}
}
