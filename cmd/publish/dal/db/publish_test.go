package db

import (
	"fmt"
	"testing"
	"time"
)

func TestAddVideo(t *testing.T) {
	Init()
	video := Video{
		ID:          1001,
		AuthorID:    1002,
		PlayURL:     "test",
		CoverURL:    "test",
		PublishTime: time.Now(),
		Title:       "test",
	}
	vid, err := CreateVideo(&video)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(vid)
}

func TestGetVideoListByLastTime(t *testing.T) {
	Init()
	lastTime := time.Now()
	videos, err := GetVideoListByLastTime(lastTime)
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestGetVideoByUserID(t *testing.T) {
	Init()
	userId := int64(1000)
	videos, err := GetVideoByUserID(userId)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestGetWorkCount(t *testing.T) {
	Init()
	count, err := GetWorkCount(1001)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(count)
}
