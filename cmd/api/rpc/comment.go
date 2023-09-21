package rpc

import (
	"context"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var commentClient commentservice.Client

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
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

	commentClient = c
}

func CommentAction(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	if commentClient == nil {
		return &comment.CommentActionResponse{}, errno.ServiceErr
	}
	resp, err := commentClient.CommentAction(ctx, req)
	if err != nil {
		return &comment.CommentActionResponse{}, err
	}

	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	if commentClient == nil {
		return &comment.CommentListResponse{}, errno.ServiceErr
	}
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return &comment.CommentListResponse{}, err
	}

	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func CommentCount(ctx context.Context, videoId int64) (int64, error) {
	if commentClient == nil {
		return 0, errno.ServiceErr
	}
	resp, err := commentClient.CommentCount(ctx, &comment.CommentCountRequest{
		VideoId: videoId,
	})
	if err != nil {
		return 0, err
	}

	return resp.Count, nil
}
