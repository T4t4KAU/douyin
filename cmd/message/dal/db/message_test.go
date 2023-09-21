package db

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestAddNewMessage(t *testing.T) {
	Init()
	msg := Message{
		ToUserId:   1002,
		FromUserId: 1003,
		Content:    "Hello World",
	}

	ok, err := AddNewMessage(context.Background(), &msg)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(ok)
}

func TestGetLatestMessageByIdPair(t *testing.T) {
	Init()
	msg, err := GetLatestMessageByIdPair(1010, 1011)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", *msg)
}

func TestGetMessageListByIdPair(t *testing.T) {
	Init()
	msgs, err := GetMessageListByIdPair(1010, 1011, time.Now())
	if err != nil {
		return
	}
	for _, msg := range msgs {
		fmt.Printf("%v\n", msg)
	}
}
