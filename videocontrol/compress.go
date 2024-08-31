package videocontrol

import (
	"fmt"
	"path/filepath"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func CompressVideo(videoPath string) {
	fileName := filepath.Base(videoPath)
	outputFile := fileName + "_H264.mp4"

	err := ffmpeg.Input(videoPath).Output(outputFile, ffmpeg.KwArgs{"c:v": "libx264"}).Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Successfully transfer to H.264 format")
}
