package modules

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	pb "viper/protos/cmds"

	"github.com/kbinani/screenshot"
)

var zeroPoint = image.Point{0, 0}

func Screenshot(req *pb.ScreenshotRequest) *pb.ScreenshotResponse {
	log.Printf("Taking a screenshot.")
	displays, err := captureDisplays()
	if err != nil {
		return &pb.ScreenshotResponse{Err: err.Error()}
	}
	display := mergeDisplays(displays)
	displayBuffer, err := encodeDisplay(display)
	if err != nil {
		return &pb.ScreenshotResponse{Err: err.Error()}
	}
	return &pb.ScreenshotResponse{Data: *displayBuffer}
}

func encodeDisplay(display *image.RGBA) (*[]byte, error) {
	displayBuffer := &bytes.Buffer{}
	err := png.Encode(displayBuffer, display)
	if err != nil {
		return nil, fmt.Errorf("Error encoding screenshot: %v", err)
	}
	displayBytes, err := io.ReadAll(displayBuffer)
	if err != nil {
		return nil, fmt.Errorf("Error reading screenshot: %v", err)
	}
	return &displayBytes, nil
}

func captureDisplays() ([]*image.RGBA, error) {
	screenCount := screenshot.NumActiveDisplays()
	var displays []*image.RGBA
	for idx := 0; idx < screenCount; idx++ {
		display, err := screenshot.CaptureDisplay(idx)
		if err != nil {
			return nil, fmt.Errorf("Error taking screenshot: %v", err)
		}
		displays = append(displays, display)
	}
	return displays, nil
}

func mergeDisplays(displays []*image.RGBA) *image.RGBA {
	mergedDisplay := displays[0]
	for _, display := range displays[1:] {
		mergedDisplay = mergeTwoImages(mergedDisplay, display)
	}
	return mergedDisplay
}

func mergeTwoImages(img1 *image.RGBA, img2 *image.RGBA) *image.RGBA {
	startPoint2 := image.Point{img1.Bounds().Dx(), 0}
	rect2 := image.Rectangle{startPoint2, startPoint2.Add(img2.Bounds().Size())}
	mergedRect := image.Rectangle{zeroPoint, rect2.Max}
	mergedImg := image.NewRGBA(mergedRect)
	draw.Draw(mergedImg, img1.Bounds(), img1, zeroPoint, draw.Src)
	draw.Draw(mergedImg, rect2, img2, zeroPoint, draw.Src)
	return mergedImg
}
