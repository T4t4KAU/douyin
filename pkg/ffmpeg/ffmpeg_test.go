package ffmpeg

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetSnapShot(t *testing.T) {
	b, err := ioutil.ReadFile("./videos/test1.mp4")
	if err != nil {
		return
	}
	shot, err := GetSnapShot(b)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	err = ioutil.WriteFile("./videos/test.jpg", shot, os.ModePerm)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}

func TestGetSnapShotByURL(t *testing.T) {
}
