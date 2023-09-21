package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

type FavoriteExistService struct {
	ctx context.Context
}

func NewFavoriteExistService(ctx context.Context) *FavoriteExistService {
	return &FavoriteExistService{
		ctx: ctx,
	}
}

func (s *FavoriteExistService) FavoriteExist(req *favorite.FavoriteExistRequest) (bool, error) {
	exist, err := rpc.PublishExist(s.ctx, req.VideoId)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errno.VideoIsNotExistErr
	}
	return db.CheckFavoriteExist(s.ctx, req.UserId, req.VideoId)
}
