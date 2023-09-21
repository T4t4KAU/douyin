package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/user"
	"sync"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) FeedAction(req *publish.FeedActionRequest) ([]*common.Video, int64, error) {
	var lastTime time.Time
	var userId int64
	var videos []*common.Video

	if req.LatestTime == nil || *req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(*req.LatestTime/1000, 0)
	}
	if req.UserId == nil {
		userId = 0
	} else {
		userId = *req.UserId
	}

	dbVideos, err := db.GetVideoListByLastTime(lastTime)
	if err != nil {
		return videos, *req.LatestTime, err
	}
	if len(dbVideos) == 0 {
		return videos, *req.LatestTime, nil
	}
	nextTime := dbVideos[len(dbVideos)-1].PublishTime.Unix()

	videoChan := make(chan common.Video, len(dbVideos))
	errChan := make(chan error, len(dbVideos))
	doneChan := make(chan struct{})

	var wg sync.WaitGroup

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
			resp, e := rpc.UserInfo(s.ctx, &user.UserInfoRequest{CurrentUserId: userId, UserId: dbVideo.AuthorID})
			if e != nil {
				errChan <- e
				return
			}
			favoritedCount, isFavorite, e := rpc.FavoriteCountOfVideo(s.ctx, userId, dbVideo.ID)
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
		return videos, nextTime, err
	default:
	}

	return videos, nextTime, nil
}
