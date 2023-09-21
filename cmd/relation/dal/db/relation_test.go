package db

import (
	"context"
	"douyin/cmd/relation/dal/cache"
	"fmt"
	"testing"
)

func TestAddNewRelation(t *testing.T) {
	Init()
	cache.Init()

	r := Relation{
		UserId:     1002,
		FollowerId: 1001,
	}

	ok, err := AddNewRelation(context.Background(), &r)
	if err != nil {
		t.Errorf("add error: %v\n", err)
		return
	}

	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}

func TestDeleteRelation(t *testing.T) {
	Init()
	cache.Init()

	r := Relation{
		UserId:     1001,
		FollowerId: 1002,
	}

	ok, err := DeleteRelation(context.Background(), &r)
	if err != nil {
		t.Errorf("delete error: %v\n", err)
		return
	}

	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}

func TestGetFollowCount(t *testing.T) {
	Init()
	cache.Init()

	ctx := context.Background()

	cnt, err := GetFollowCount(ctx, 1001)
	if err != nil {
		t.Errorf("get error: %v", err)
		return
	}
	fmt.Println(cnt)
}

func TestGetFollowerCount(t *testing.T) {
	Init()
	cache.Init()

	ctx := context.Background()

	cnt, err := GetFollowerCount(ctx, 1001)
	if err != nil {
		t.Errorf("get error: %v", err)
		return
	}
	fmt.Println(cnt)
}

func TestGetFollowIdList(t *testing.T) {
	Init()
	cache.Init()

	ctx := context.Background()

	followerList, err := GetFollowIdList(ctx, 1002)
	if err != nil {
		t.Errorf("get error: %v", err)
		return
	}

	for _, id := range followerList {
		fmt.Println(id)
	}
}

func TestGetFollowerIdList(t *testing.T) {
	Init()
	cache.Init()

	ctx := context.Background()

	followerList, err := GetFollowerIdList(ctx, 1001)
	if err != nil {
		t.Errorf("get error: %v", err)
		return
	}

	for _, id := range followerList {
		fmt.Println(id)
	}
}

func TestCheckRelationFollowExist(t *testing.T) {
	Init()
	cache.Init()

	exist, err := CheckRelationFollowExist(context.Background(), 1001, 1002)
	if err != nil {
		t.Errorf(err.Error())
	}
	println(exist)
}
