package rpc

import (
	"context"
	"douyin/kitex_gen/publish"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPublishAction(t *testing.T) {
	InitRPC()

	b, _ := ioutil.ReadFile("./videos/hwx.mp4")
	resp, err := PublishAction(context.Background(), &publish.PublishActionRequest{
		UserId: 1001,
		Data:   b,
		Title:  "test",
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%#v\n", resp)
}

func TestPublishCount(t *testing.T) {
	InitRPC()

	count, err := PublishCount(context.Background(), 1001)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(count)
}

func TestPublishExist(t *testing.T) {
	InitRPC()

	exist, err := PublishExist(context.Background(), 1000)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(exist)
}

func TestPublishList(t *testing.T) {
	InitRPC()

	resp, err := PublishList(context.Background(), &publish.PublishListRequest{
		CurrentUserId: 1010,
		UserId:        1010,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, item := range resp.VideoList {
		fmt.Printf("%#v\n", *item)
	}
}

func TestPublishInfo(t *testing.T) {
	InitRPC()

	video, err := PublishInfo(context.Background(), 1018)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("%#v\n", *video)
}

func TestPublishVideoList(t *testing.T) {
	InitRPC()

	videos, err := PublishVideoList(context.Background(), []int64{1024, 1027})

	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", *video)
	}
}

func TestFeedAction(t *testing.T) {
	InitRPC()

	resp, err := FeedAction(context.Background(), &publish.FeedActionRequest{})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for _, video := range resp.VideoList {
		fmt.Printf("%#v\n", *video)
	}
}
