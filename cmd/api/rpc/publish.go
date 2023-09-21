package rpc

import (
	"context"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/publish/publishservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var publishClient publishservice.Client

func initPublishRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServiceName,
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),               // mux
		client.WithRPCTimeout(3*time.Second),      // rpc timeout
		client.WithConnectTimeout(10*time.Minute), // conn timeout
		client.WithRPCTimeout(10*time.Minute),
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}

	publishClient = c
}

func PublishAction(ctx context.Context, req *publish.PublishActionRequest) (*publish.PublishActionResponse, error) {
	resp, err := publishClient.PublishAction(ctx, req)
	if err != nil {
		return &publish.PublishActionResponse{}, err
	}

	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func PublishCount(ctx context.Context, userId int64) (int64, error) {
	req := &publish.PublishCountRequest{
		UserId: userId,
	}
	resp, err := publishClient.PublishCount(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.WorkCount, err
}

func PublishList(ctx context.Context, req *publish.PublishListRequest) (*publish.PublishListResponse, error) {
	resp, err := publishClient.PublishList(ctx, req)
	if resp == nil {
		return &publish.PublishListResponse{}, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, err
}

func PublishExist(ctx context.Context, videoId int64) (bool, error) {
	if publishClient == nil {
		return false, errno.ServiceErr
	}
	resp, err := publishClient.PublishExist(ctx, &publish.PublishExistRequest{
		VideoId: videoId,
	})
	return resp.Exist, err
}

func PublishInfo(ctx context.Context, videoId int64) (*common.Video, error) {
	if publishClient == nil {
		return &common.Video{}, errno.ServiceErr
	}
	resp, err := publishClient.PublishInfo(ctx, &publish.PublishInfoRequest{
		VideoId: videoId,
	})
	if err != nil {
		return &common.Video{}, err
	}
	return resp.Video, nil
}

func PublishVideoList(ctx context.Context, videoIds []int64) ([]*common.Video, error) {
	if publishClient == nil {
		return []*common.Video{}, errno.ServiceErr
	}

	req := publish.PublishVideoListRequest{
		VideoIds: videoIds,
	}
	resp, err := publishClient.PublishVideoList(ctx, &req)
	if err != nil {
		return []*common.Video{}, err
	}
	return resp.VideoList, nil
}

func FeedAction(ctx context.Context, req *publish.FeedActionRequest) (*publish.FeedActionResponse, error) {
	if publishClient == nil {
		return &publish.FeedActionResponse{}, errno.ServiceErr
	}
	resp, err := publishClient.FeedAction(ctx, req)
	if err != nil {
		return &publish.FeedActionResponse{}, err
	}
	return resp, nil
}
