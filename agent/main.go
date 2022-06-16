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
			echoCmd := cmd.GetEchoCommandRequest()
			conn.Write([]byte(echoCmd.Text))
		}
	}
}
