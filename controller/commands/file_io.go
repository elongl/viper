package commands

import (
	"context"
	"fmt"
	"log"
	"viper/controller/agents"
	pb "viper/protos/cmds"
)

func (s *AgentManagerServer) DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	log.Printf("[%d] Downloading file '%s'.", req.AgentId, req.Path)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.DownloadFile(req)
	if err != nil {
		return nil, err
	}
	if resp.Err != "" {
		return nil, fmt.Errorf("Failed to download file: %v", resp.Err)
	}
	log.Print("Downloaded file.")
	return resp, nil
}

func (s *AgentManagerServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	log.Printf("[%d] Uploading file '%s'.", req.AgentId, req.Path)
	agent, err := agents.GetAgent(req.AgentId)
	if err != nil {
		return nil, err
	}
	resp, err := agent.UploadFile(req)
	if err != nil {
		return nil, err
	}
	if resp.Err != "" {
		return nil, fmt.Errorf("Failed to upload file: %v", resp.Err)
	}
	log.Print("Uploaded file.")
	return resp, nil
}
