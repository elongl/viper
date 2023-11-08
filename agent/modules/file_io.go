package modules

import (
	pb "agent/protos/cmds"
	"log"
	"os"
)

func UploadFileToController(req *pb.DownloadFileRequest) *pb.DownloadFileResponse {
	log.Printf("uploading file to controller from '%s'", req.Path)
	data, err := os.ReadFile(req.Path)
	if err != nil {
		log.Printf("failed to download file: %v", err)
		return &pb.DownloadFileResponse{Err: err.Error()}
	}
	return &pb.DownloadFileResponse{Data: data}
}

func DownloadFileFromController(req *pb.UploadFileRequest) *pb.UploadFileResponse {
	log.Printf("downloading file from controller to '%s'", req.Path)
	err := os.WriteFile(req.Path, req.Data, 0644)
	if err != nil {
		log.Printf("failed to upload file: %v", err)
		return &pb.UploadFileResponse{Err: err.Error()}
	}
	return &pb.UploadFileResponse{}
}
