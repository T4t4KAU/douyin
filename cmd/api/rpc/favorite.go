package rpc

import (
	"context"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoriteservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
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

	favoriteClient = c
}

func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	if favoriteClient == nil {
		return &favorite.FavoriteActionResponse{}, errno.ServiceErr
	}
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return &favorite.FavoriteActionResponse{}, err
	}

	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteCount(ctx context.Context, userId int64) (int64, int64, error) {
	if favoriteClient == nil {
		return 0, 0, errno.ServiceErr
	}
	req := &favorite.FavoriteCountRequest{
		UserId: userId,
	}
	resp, err := favoriteClient.FavoriteCount(ctx, req)
	if err != nil {
		return 0, 0, err
	}
	return resp.FavoriteCount, resp.FavoritedCount, err
}

func FavoriteExist(ctx context.Context, userId int64, videoId int64) (bool, error) {
	if favoriteClient == nil {
		return false, errno.ServiceErr
	}
	req := &favorite.FavoriteExistRequest{
		UserId:  userId,
		VideoId: videoId,
	}
	resp, err := favoriteClient.FavoriteExist(ctx, req)
	return resp.Exist, err
}

func FavoriteCountOfVideo(ctx context.Context, userId int64, videoId int64) (int64, bool, error) {
	if favoriteClient == nil {
		return -1, false, errno.ServiceErr
	}
	req := &favorite.FavoriteCountOfVideoRequest{
		UserId:  userId,
		VideoId: videoId,
	}
	resp, err := favoriteClient.FavoriteCountOfVideo(ctx, req)
	if err != nil {
		return -1, false, err
	}
	return resp.FavoritedCount, resp.IsFavorite, nil
}

func FavoriteVideoList(ctx context.Context, req *favorite.FavoriteVideoListRequest) (*favorite.FavoriteVideoListResponse, error) {
	if favoriteClient == nil {
		return &favorite.FavoriteVideoListResponse{}, nil
	}
	resp, err := favoriteClient.FavoriteVideoList(ctx, req)
	if err != nil {
		return &favorite.FavoriteVideoListResponse{}, err
	}
	return resp, nil
}
