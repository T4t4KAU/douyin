package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/mq"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"strconv"
	"strings"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{
		ctx: ctx,
	}
}

func (s *CommentActionService) CommentAction(req *comment.CommentActionRequest) (*comment.Comment, error) {
	c := new(comment.Comment)

	if req.ActionType == 1 {
		dbComment := &db.Comment{
			UserId:      req.UserId,
			VideoId:     req.VideoId,
			CommentText: req.CommentText,
		}

		exist, err := rpc.PublishExist(s.ctx, req.VideoId)
		if err != nil {
			return c, err
		}
		if !exist {
			return c, errno.VideoIsNotExistErr
		}

		sb := strings.Builder{}
		sb.WriteString(strconv.Itoa(int(req.UserId)))
		sb.WriteString("&")
		sb.WriteString(strconv.Itoa(int(req.VideoId)))
		sb.WriteString("&")
		sb.WriteString(req.CommentText)

		err = mq.AddActor.Publish(s.ctx, sb.String())
		if err != nil {
			return c, err
		}

		c.Id = dbComment.ID
		c.CreateDate = dbComment.CreatedAt.Format("01-02")
		c.Content = &dbComment.CommentText
		resp, err := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
			CurrentUserId: 0,
			UserId:        req.UserId,
		})
		if err != nil {
			return c, err
		}
		c.User = resp.User
	} else {
		err := db.DeleteCommentById(s.ctx, c.Id)
		if err != nil {
			return c, err
		}
	}
	return c, nil
}
