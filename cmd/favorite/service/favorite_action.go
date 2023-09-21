package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/dal/mq"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"strconv"
	"strings"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	exist, err := rpc.PublishExist(s.ctx, req.VideoId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.VideoIsNotExistErr
	}

	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(req.UserId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.VideoId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.ActionType)))

	return mq.AddActor.Publish(s.ctx, sb.String())
}
