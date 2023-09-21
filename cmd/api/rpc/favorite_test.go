package rpc

import (
	"context"
	"douyin/kitex_gen/favorite"
	"fmt"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	InitRPC()

	resp, err := FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
		UserId:     1001,
		VideoId:    1003,
		ActionType: 1,
	})

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Printf("%#v\n", resp)
}

func TestFavoriteCount(t *testing.T) {
	InitRPC()

	c1, c2, err := FavoriteCount(context.Background(), 1001)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(c1, c2)
}

func TestFavoriteExist(t *testing.T) {
	InitRPC()

	exist, err := FavoriteExist(context.Background(), 1010, 1018)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(exist)
}

func TestFavoriteCountOfVideo(t *testing.T) {
	InitRPC()

	count, exist, err := FavoriteCountOfVideo(context.Background(), 1020, 1023)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(count, exist)
}

func TestFavoriteVideoList(t *testing.T) {
	InitRPC()

	resp, err := FavoriteVideoList(context.Background(), &favorite.FavoriteVideoListRequest{
		UserId:   1020,
		ToUserId: 1020,
	})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for _, video := range resp.VideoList {
		fmt.Printf("%#v\n", *video)
	}
}
