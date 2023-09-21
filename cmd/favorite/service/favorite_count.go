package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

type FavoriteCountService struct {
	ctx context.Context
}

func NewFavoriteCountService(ctx context.Context) *FavoriteCountService {
	return &FavoriteCountService{
		ctx: ctx,
	}
}

func (s *FavoriteCountService) FavoriteCount(req *favorite.FavoriteCountRequest) (int64, int64, error) {
	exist, err := rpc.UserExist(s.ctx, req.UserId)
	if err != nil {
		return 0, 0, err
	}
	if !exist {
		return 0, 0, errno.UserIsNotExistErr
	}

	favoriteCount, err := db.GetUserFavoriteCountById(s.ctx, req.UserId)
	favoritedCount, err := db.GetTotalFavoritedbCountByAuthorID(s.ctx, req.UserId)
	return favoriteCount, favoritedCount, err
}
