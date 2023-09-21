package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

type FavoriteVideoListService struct {
	ctx context.Context
}

func NewFavoriteVideoListService(ctx context.Context) *FavoriteVideoListService {
	return &FavoriteVideoListService{
		ctx: ctx,
	}
}

func (s *FavoriteVideoListService) FavoriteVideoList(req *favorite.FavoriteVideoListRequest) ([]*common.Video, error) {
	var videos []*common.Video

	exist, err := rpc.UserExist(s.ctx, req.ToUserId)
	if err != nil {
		return videos, err
	}
	if !exist {
		return videos, errno.UserIsNotExistErr
	}

	videoIds, err := db.GetUserFavoriteIdList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos, err = rpc.PublishVideoList(s.ctx, videoIds)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(videos); i++ {
		videos[i].FavoriteCount, err = db.GetVideoFavoritedCount(s.ctx, videos[i].Id)
		if err != nil {
			return nil, err
		}
	}

	return videos, nil
}
