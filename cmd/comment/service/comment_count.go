package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/kitex_gen/comment"
)

type CommentCountService struct {
	ctx context.Context
}

func NewCommentCountService(ctx context.Context) *CommentCountService {
	return &CommentCountService{
		ctx: ctx,
	}
}

func (s *CommentCountService) CommentCount(req *comment.CommentCountRequest) (int64, error) {
	return db.GetCommentCountByVideoID(s.ctx, req.VideoId)
}
