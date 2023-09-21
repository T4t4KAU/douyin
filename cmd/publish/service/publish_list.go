package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/user"
	"sync"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

func (s *PublishListService) PublishList(req *publish.PublishListRequest) ([]*common.Video, error) {
	var videos []*common.Video

	// 获取用户发布视频
	dbVideos, err := db.GetVideoByUserID(req.UserId)
	if err != nil {
		return videos, err
	}

	// 获取用户信息
	resp, err := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
		CurrentUserId: req.CurrentUserId,
		UserId:        req.UserId,
	})
	if err != nil {
		return videos, err
	}

	var wg sync.WaitGroup

	videoChan := make(chan common.Video, len(dbVideos))
	errChan := make(chan error, len(dbVideos))
	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case v := <-videoChan:
				videos = append(videos, &v)
			case <-doneChan:
				return
			}
		}
	}()

	for _, v := range dbVideos {
		wg.Add(1)
		go func(dbVideo db.Video) {
			defer wg.Done()
			favoritedCount, isFavorite, e := rpc.FavoriteCountOfVideo(s.ctx, req.UserId, dbVideo.ID)
			if e != nil {
				errChan <- e
				return
			}
			commentCount, e := rpc.CommentCount(s.ctx, dbVideo.ID)
			if e != nil {
				errChan <- e
				return
			}
			videoChan <- common.Video{
				Id:            dbVideo.ID,
				Author:        resp.User,
				PlayUrl:       dbVideo.PlayURL,
				CoverUrl:      dbVideo.CoverURL,
				FavoriteCount: favoritedCount,
				IsFavorite:    isFavorite,
				Title:         dbVideo.Title,
				CommentCount:  commentCount,
			}
		}(*v)
	}

	wg.Wait()
	doneChan <- struct{}{}

	select {
	case err = <-errChan:
		return videos, err
	default:
	}

	return videos, nil
}
