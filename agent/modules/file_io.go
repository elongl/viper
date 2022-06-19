package modules

import (
	"log"
	"os"
	pb "viper/protos/cmds"
)

func UploadFileToController(req *pb.DownloadFileRequest) *pb.DownloadFileResponse {
	log.Printf("Uploading file to controller from '%s' on agent '%d'.", req.Path, req.AgentId)
	data, err := os.ReadFile(req.Path)
	if err != nil {
		log.Printf("Failed to download file: %v", err)
		return &pb.DownloadFileResponse{Err: err.Error()}
	}
	return &pb.DownloadFileResponse{Data: data}
}

func DownloadFileFromController(req *pb.UploadFileRequest) *pb.UploadFileResponse {
	log.Printf("Downloading file from controller to '%s' on agent '%d'.", req.Path, req.AgentId)
	err := os.WriteFile(req.Path, req.Data, 0644)
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		return &pb.UploadFileResponse{Err: err.Error()}
	}
	return &pb.UploadFileResponse{}
}
