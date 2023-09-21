package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var relationClient relationservice.Client

func initRelationRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := relationservice.NewClient(
		constants.RelationServiceName,
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}

	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {
	if relationClient == nil {
		return &relation.RelationActionResponse{}, errno.ServiceErr
	}
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return &relation.RelationActionResponse{}, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func RelationCount(ctx context.Context, userId int64) (int64, int64, error) {
	if relationClient == nil {
		return 0, 0, errno.ServiceErr
	}
	resp, err := relationClient.RelationCount(ctx, &relation.RelationCountRequest{
		UserId: userId,
	})
	return resp.FollowCount, resp.FollowerCount, err
}

func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (*relation.RelationFollowListResponse, error) {
	if relationClient == nil {
		return &relation.RelationFollowListResponse{}, errno.ServiceErr
	}
	resp, err := relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return &relation.RelationFollowListResponse{}, err
	}
	return resp, nil
}

func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (*relation.RelationFollowerListResponse, error) {
	if relationClient == nil {
		return &relation.RelationFollowerListResponse{}, errno.ServiceErr
	}
	resp, err := relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return &relation.RelationFollowerListResponse{}, err
	}
	return resp, nil
}

func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (*relation.RelationFriendListResponse, error) {
	if relationClient == nil {
		return &relation.RelationFriendListResponse{}, errno.ServiceErr
	}
	return relationClient.RelationFriendList(ctx, req)
}

func RelationExist(ctx context.Context, currentUserId int64, userId int64) (bool, bool, error) {
	if relationClient == nil {
		return false, false, errno.ServiceErr
	}
	resp, err := relationClient.RelationExist(ctx, &relation.RelationExistRequest{
		CurrentUserId: currentUserId,
		UserId:        userId,
	})
	if err != nil {
		return false, false, err
	}
	return resp.FollowExist, resp.FollowedExist, nil
}
