package rpc

import (
	"context"
	"douyin/kitex_gen/comment"
	"fmt"
	"testing"
)

func TestCommentAction(t *testing.T) {
	InitRPC()

	resp, err := CommentAction(context.Background(), &comment.CommentActionRequest{
		UserId:      1001,
		VideoId:     1002,
		ActionType:  1,
		CommentText: "hello",
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%#v\n", resp)
}

func TestDeleteComment(t *testing.T) {
	InitRPC()

	commentId := int64(2)
	resp, err := CommentAction(context.Background(), &comment.CommentActionRequest{
		UserId:      1001,
		VideoId:     1003,
		ActionType:  2,
		CommentText: "hello",
		CommentId:   &commentId,
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%#v\n", resp.Comment)
}

func TestCommentList(t *testing.T) {
	InitRPC()

	resp, err := CommentList(context.Background(), &comment.CommentListRequest{
		UserId:  1021,
		VideoId: 1023,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("%#v\n", resp.CommentList)
}

func TestCommentCount(t *testing.T) {
	InitRPC()

	count, err := CommentCount(context.Background(), 1018)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("%#v\n", count)
}
