package service

import (
	"context"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/publish"
)

type PublishCountService struct {
	ctx context.Context
}

func NewPublishCountService(ctx context.Context) *PublishCountService {
	return &PublishCountService{
		ctx: ctx,
	}
}

func (s *PublishCountService) PublishCount(req *publish.PublishCountRequest) (int64, error) {
	return db.GetWorkCount(req.UserId)
}
