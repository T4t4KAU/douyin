package db

import (
	"context"
	"fmt"
	"testing"
)

func TestAddNewComment(t *testing.T) {
	Init()
	comment := &Comment{
		UserId:      1010,
		VideoId:     1018,
		CommentText: "video comment test",
	}
	err := AddNewComment(context.Background(), comment)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func TestDeleteComment(t *testing.T) {
	Init()
	err := DeleteCommentById(context.Background(), 1020)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func TestGetCommentListByVideoID(t *testing.T) {
	Init()
	comments, err := GetCommentListByVideoID(context.Background(), 1023)
	if err != nil {
		return
	}
	for _, c := range comments {
		fmt.Printf("%v %v %v %v %v\n", c.ID, c.UserId, c.VideoId,
			c.CommentText, c.CreatedAt.Format("01-02"))
	}
}
