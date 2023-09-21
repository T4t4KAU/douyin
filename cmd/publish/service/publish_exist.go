package service

import (
	"context"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/publish"
)

type PublishExistService struct {
	ctx context.Context
}

func NewPublishExistService(ctx context.Context) *PublishCountService {
	return &PublishCountService{
		ctx: ctx,
	}
}

func (s *PublishCountService) PublishExist(req *publish.PublishExistRequest) (bool, error) {
	return db.CheckVideoExistById(req.VideoId)
}
