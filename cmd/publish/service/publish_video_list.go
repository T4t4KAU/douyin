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

type PublishVideoListService struct {
	ctx context.Context
}

func NewPublishVideoListService(ctx context.Context) *PublishVideoListService {
	return &PublishVideoListService{
		ctx: ctx,
	}
}

func (s *PublishVideoListService) PublishVideoList(req *publish.PublishVideoListRequest) ([]*common.Video, error) {
	var videos []*common.Video

	dbVideos, err := db.GetVideoListByVideoIDList(req.VideoIds)
	if err != nil {
		return videos, err
	}

	num := len(dbVideos)
	videoChan := make(chan common.Video, num)
	errChan := make(chan error, num)
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
		go func(video *db.Video) {
			defer wg.Done()
			commentCount, e := rpc.CommentCount(s.ctx, video.ID)
			if e != nil {
				errChan <- e
				return
			}

			resp, e := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
				CurrentUserId: 0,
				UserId:        video.AuthorID,
			})
			if e != nil {
				errChan <- e
				return
			}

			videoChan <- common.Video{
				Id:            video.ID,
				PlayUrl:       video.PlayURL,
				CoverUrl:      video.CoverURL,
				Title:         video.Title,
				CommentCount:  commentCount,
				FavoriteCount: 0,
				IsFavorite:    true,
				Author:        resp.User,
			}

		}(v)
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
