package agents

import (
	"encoding/binary"
	"errors"
	"log"
	"net"

	"google.golang.org/protobuf/proto"

	pb "viper/protos/cmds"
)

type Agent struct {
	Hostname string
	Alive    bool
	conn     net.Conn
}

var (
	agents []Agent
)

func GetAgent(hostname string) (*Agent, error) {
	for _, agent := range agents {
		if agent.Hostname == hostname {
			return &agent, nil
		}
	}
	return nil, errors.New("Agent not found")
}

func InitAgent(conn net.Conn) {
	log.Printf("Initializing agent: %v", conn.RemoteAddr())
	agent := Agent{conn: conn, Hostname: "egk", Alive: true}
	agents = append(agents, agent)
}

func (agent *Agent) Read(resp proto.Message) error {
	var respSize int64
	err := binary.Read(agent.conn, binary.LittleEndian, &respSize)
	if err != nil {
		return err
	}
	respBuffer := make([]byte, respSize)
	_, err = agent.conn.Read(respBuffer)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(respBuffer, resp)
	if err != nil {
		return err
	}
	return nil
}

func (agent *Agent) Write(cmdReq *pb.CommandRequest) error {
	cmdBuffer, err := proto.Marshal(cmdReq)
	if err != nil {
		return err
	}
	cmdBufferLen := int64(len(cmdBuffer))
	err = binary.Write(agent.conn, binary.LittleEndian, &cmdBufferLen)
	if err != nil {
		return err
	}
	_, err = agent.conn.Write(cmdBuffer)
	if err != nil {
		return err
	}
	return nil
}

func (agent *Agent) Close() {
	log.Printf("Closing agent: %v", agent.Hostname)
	agent.conn.Close()
	agent.Alive = false
}
