package agents

import (
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/proto"

	pb "viper/protos/cmds"
)

type Agent struct {
	id          int64
	alive       bool
	conn        net.Conn
	connectTime time.Time
}

const (
	ERR_AGENT_NO_LONGER_CONNECTED = "Agent no longer connected."
	ERR_AGENT_NOT_FOUND           = "Agent not found."
)

var (
	agents = make(map[int64]*Agent)
)

func GetAgent(id int64) (*Agent, error) {
	agent := agents[id]
	if agent == nil {
		return nil, errors.New(ERR_AGENT_NOT_FOUND)
	}
	if !agent.alive {
		return nil, errors.New(ERR_AGENT_NO_LONGER_CONNECTED)
	}
	return agent, nil
}

func InitAgent(conn net.Conn) {
	log.Printf("Initializing agent: %v", conn.RemoteAddr())
	agentId := int64(len(agents))
	agent := &Agent{conn: conn, id: agentId, alive: true, connectTime: time.Now()}
	agents[agentId] = agent
}

func (agent *Agent) RunEchoCommand(req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	cmdReq := &pb.CommandRequest{Type: pb.ECHO_CMD_TYPE, Req: &pb.CommandRequest_EchoCommandRequest{EchoCommandRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.EchoCommandResponse{}
	err = agent.read(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (agent *Agent) read(resp proto.Message) error {
	var respSize int64
	err := binary.Read(agent.conn, binary.LittleEndian, &respSize)
	if err == io.EOF {
		agent.Close()
		return errors.New(ERR_AGENT_NO_LONGER_CONNECTED)
	}
	if err != nil {
		return err
	}
	respBuffer := make([]byte, respSize)
	_, err = agent.conn.Read(respBuffer)
	if err == io.EOF {
		agent.Close()
		return errors.New(ERR_AGENT_NO_LONGER_CONNECTED)
	}
	if err != nil {
		return err
	}
	err = proto.Unmarshal(respBuffer, resp)
	if err != nil {
		return err
	}
	return nil
}

func (agent *Agent) write(cmdReq *pb.CommandRequest) error {
	cmdBuffer, err := proto.Marshal(cmdReq)
	if err != nil {
		return err
	}
	cmdBufferLen := int64(len(cmdBuffer))
	err = binary.Write(agent.conn, binary.LittleEndian, &cmdBufferLen)
	if err == io.EOF {
		agent.Close()
		return errors.New(ERR_AGENT_NO_LONGER_CONNECTED)
	}
	if err != nil {
		return err
	}
	_, err = agent.conn.Write(cmdBuffer)
	if err == io.EOF {
		agent.Close()
		return errors.New(ERR_AGENT_NO_LONGER_CONNECTED)
	}
	if err != nil {
		return err
	}
	return nil
}

func (agent *Agent) Close() {
	log.Printf("Agent '%d' is no longer connected.", agent.id)
	agent.conn.Close()
	agent.alive = false
}
