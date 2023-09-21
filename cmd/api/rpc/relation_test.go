package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"fmt"
	"testing"
)

func TestRelationAction(t *testing.T) {
	InitRPC()

	resp, err := RelationAction(context.Background(), &relation.RelationActionRequest{
		CurrentUserId: 1023,
		ToUserId:      1020,
		ActionType:    1,
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%#v\n", resp)
}

func TestRelationCount(t *testing.T) {
	InitRPC()

	c1, c2, err := RelationCount(context.Background(), 1010)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(c1, c2)
}

func TestRelationExist(t *testing.T) {
	InitRPC()

	e1, e2, err := RelationExist(context.Background(), 1010, 1011)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(e1, e2)
}

func TestRelationFollowList(t *testing.T) {
	InitRPC()

	resp, err := RelationFollowList(context.Background(), &relation.RelationFollowListRequest{
		UserId:        1010,
		CurrentUserId: 1010,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, u := range resp.UserList {
		fmt.Printf("%#v\n", u)
	}
}

func TestRelationFollowerList(t *testing.T) {
	InitRPC()

	resp, err := RelationFollowerList(context.Background(), &relation.RelationFollowerListRequest{
		UserId:        1010,
		CurrentUserId: 1011,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, u := range resp.UserList {
		fmt.Printf("%#v\n", u)
	}
}

func TestRelationFriendList(t *testing.T) {
	InitRPC()

	resp, err := RelationFriendList(context.Background(), &relation.RelationFriendListRequest{
		UserId:        1023,
		CurrentUserId: 1020,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, u := range resp.UserList {
		fmt.Printf("%#v\n", u)
	}
}
