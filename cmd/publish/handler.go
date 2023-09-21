package main

import (
	"context"
	"douyin/cmd/publish/pkg"
	"douyin/cmd/publish/service"
	publish "douyin/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp = new(publish.PublishActionResponse)

	_, err = service.NewPublishActionService(ctx).PublishAction(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	return
}

// PublishCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishCount(ctx context.Context, req *publish.PublishCountRequest) (resp *publish.PublishCountResponse, err error) {
	resp = new(publish.PublishCountResponse)

	count, err := service.NewPublishCountService(ctx).PublishCount(req)
	resp.WorkCount = count

	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp = new(publish.PublishListResponse)

	videos, err := service.NewPublishListService(ctx).PublishList(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.VideoList = videos

	return
}

// PublishExist implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishExist(ctx context.Context, req *publish.PublishExistRequest) (resp *publish.PublishExistResponse, err error) {
	resp = new(publish.PublishExistResponse)

	exist, err := service.NewPublishExistService(ctx).PublishExist(req)
	resp.Exist = exist

	return
}

// PublishInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishInfo(ctx context.Context, req *publish.PublishInfoRequest) (resp *publish.PublishInfoResponse, err error) {
	resp = new(publish.PublishInfoResponse)

	resp.Video, err = service.NewPublishInfoService(ctx).PublishInfo(req)

	return
}

// PublishVideoList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideoList(ctx context.Context, req *publish.PublishVideoListRequest) (resp *publish.PublishVideoListResponse, err error) {
	resp = new(publish.PublishVideoListResponse)

	resp.VideoList, err = service.NewPublishVideoListService(ctx).PublishVideoList(req)

	return
}

// PublishListByLastTime implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishListByLastTime(ctx context.Context, req *publish.PublishListByLastTimeRequest) (resp *publish.PublishListByLastTimeResponse, err error) {
	// TODO: Your code here...
	return
}

// FeedAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) FeedAction(ctx context.Context, req *publish.FeedActionRequest) (resp *publish.FeedActionResponse, err error) {
	resp = new(publish.FeedActionResponse)

	videos, nextTime, err := service.NewFeedService(ctx).FeedAction(req)
	resp.NextTime = &nextTime
	resp.VideoList = videos

	return
}
