package ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image/jpeg"
	"os"
	"os/exec"
)

// GetSnapShotByPath 获取视频截图
func GetSnapShotByPath(path string) (*bytes.Buffer, error) {
	buff := bytes.NewBuffer(nil)
	err := ffmpeg.Input(path).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buff).Run()
	return buff, err
}

func GetSnapShot(video []byte) ([]byte, error) {
	inputBuffer := bytes.NewBuffer(video)
	outputBuffer := bytes.NewBuffer(nil)

	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vf", `select=gte(n\,0)`,
		"-vframes", "1", "-f", "image2", "pipe:1")
	cmd.Stdin, cmd.Stdout = inputBuffer, outputBuffer

	err := cmd.Run()
	return outputBuffer.Bytes(), err
}

func GetSnapShotByURL(url string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(url).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	img, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	err = jpeg.Encode(buff, img, nil)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
