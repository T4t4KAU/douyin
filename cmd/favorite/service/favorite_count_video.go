package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

type FavoriteCountOfVideoService struct {
	ctx context.Context
}

func NewFavoriteCountOfVideoService(ctx context.Context) *FavoriteCountOfVideoService {
	return &FavoriteCountOfVideoService{
		ctx: ctx,
	}
}

func (s *FavoriteCountOfVideoService) FavoriteCountOfVideo(req *favorite.FavoriteCountOfVideoRequest) (int64, bool, error) {
	exist, err := rpc.PublishExist(s.ctx, req.VideoId)
	if err != nil {
		return -1, false, err
	}
	if !exist {
		return -1, false, errno.VideoIsNotExistErr
	}
	count, err := db.GetVideoFavoritedCount(s.ctx, req.VideoId)
	if err != nil {
		return -1, false, err
	}
	exist, err = db.CheckFavoriteExist(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		return -1, false, err
	}
	return count, exist, err
}
