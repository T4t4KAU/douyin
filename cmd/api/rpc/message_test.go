package rpc

import (
	"context"
	"douyin/kitex_gen/message"
	"fmt"
	"testing"
)

func TestMessageAction(t *testing.T) {
	InitRPC()

	resp, err := MessageAction(context.Background(), &message.MessageActionRequest{
		UserId:     1001,
		ToUserId:   1003,
		ActionType: 1,
		Content:    "douyin test",
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%#v\n", resp)
}
