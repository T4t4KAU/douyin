package db

import (
	"context"
	"douyin/cmd/favorite/dal/cache"
	"fmt"
	"testing"
)

func TestAddNewFavorite(t *testing.T) {
	Init()
	cache.Init()

	msg := Favorites{
		UserId:  1002,
		VideoId: 1004,
	}

	ok, err := AddNewFavorite(context.Background(), &msg)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(ok)
}

func TestDeleteFavorite(t *testing.T) {
	Init()
	cache.Init()

	msg := Favorites{
		UserId:  1010,
		VideoId: 1018,
	}

	ok, err := DeleteFavorite(context.Background(), &msg)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(ok)
}

func TestGetUserFavoriteCountById(t *testing.T) {
	Init()
	cache.Init()

	count, err := GetUserFavoriteCountById(context.Background(), 1002)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(count)
}

func TestGetVideoFavoritedCount(t *testing.T) {
	Init()
	cache.Init()

	count, err := GetVideoFavoritedCount(context.Background(), 1003)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(count)
}

func TestUserFavoriteIdList(t *testing.T) {
	Init()
	cache.Init()

	videos, err := getUserFavoriteIdList(1010)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("%#v\n", videos)
}

func TestCheckFavoriteExist(t *testing.T) {
	Init()
	cache.Init()

	msg := Favorites{
		UserId:  1010,
		VideoId: 1018,
	}

	ok, err := CheckFavoriteExist(context.Background(), msg.UserId, msg.VideoId)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(ok)
}

func TestGetTotalFavoritedbCountByAuthorID(t *testing.T) {
	Init()
	cache.Init()

	count, err := GetTotalFavoritedbCountByAuthorID(context.Background(), 1010)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	println(count)
}
