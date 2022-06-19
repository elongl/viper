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
	Id          int64
	Alive       bool
	conn        net.Conn
	ConnectTime time.Time
}

const (
	ERR_AGENT_NO_LONGER_CONNECTED = "Agent no longer connected."
	ERR_AGENT_NOT_FOUND           = "Agent not found."
)

var (
	Agents = make(map[int64]*Agent)
)

func GetAgent(id int64) (*Agent, error) {
	agent := Agents[id]
	if agent == nil {
		return nil, errors.New(ERR_AGENT_NOT_FOUND)
	}
	if !agent.Alive {
		return nil, errors.New(ERR_AGENT_NO_LONGER_CONNECTED)
	}
	return agent, nil
}

func InitAgent(conn net.Conn) {
	agentId := int64(len(Agents))
	log.Printf("Initializing agent (%d) @ %v", agentId, conn.RemoteAddr())
	agent := &Agent{conn: conn, Id: agentId, Alive: true, ConnectTime: time.Now()}
	Agents[agentId] = agent
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

func (agent *Agent) RunShellCommand(req *pb.ShellCommandRequest) (*pb.ShellCommandResponse, error) {
	cmdReq := &pb.CommandRequest{Type: pb.SHELL_CMD_TYPE, Req: &pb.CommandRequest_ShellCommandRequest{ShellCommandRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.ShellCommandResponse{}
	err = agent.read(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (agent *Agent) DownloadFile(req *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	cmdReq := &pb.CommandRequest{Type: pb.DOWNLOAD_FILE_CMD_TYPE, Req: &pb.CommandRequest_DownloadFileRequest{DownloadFileRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.DownloadFileResponse{}
	err = agent.read(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (agent *Agent) UploadFile(req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	cmdReq := &pb.CommandRequest{Type: pb.UPLOAD_FILE_CMD_TYPE, Req: &pb.CommandRequest_UploadFileRequest{UploadFileRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.UploadFileResponse{}
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
	log.Printf("Agent '%d' is no longer connected.", agent.Id)
	agent.conn.Close()
	agent.Alive = false
}
