package rpc

import (
	"context"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messageservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var messageClient messageservice.Client

func initMessageRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		constants.MessageServiceName,
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

	messageClient = c
}

func MessageAction(ctx context.Context, req *message.MessageActionRequest) (*message.MessageActionResponse, error) {
	if messageClient == nil {
		return &message.MessageActionResponse{}, errno.ServiceErr
	}
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return &message.MessageActionResponse{}, err
	}

	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func MessageList(ctx context.Context, req *message.MessageListRequest) (*message.MessageListResponse, error) {
	if messageClient == nil {
		return &message.MessageListResponse{}, errno.ServiceErr
	}
	resp, err := messageClient.MessageList(ctx, req)
	if err != nil {
		return &message.MessageListResponse{}, err
	}

	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
