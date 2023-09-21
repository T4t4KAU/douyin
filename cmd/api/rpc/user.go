package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var userClient userservice.Client

// 初始化用户RPC
func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return &user.UserRegisterResponse{}, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (int64, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return int64(0), err
	}
	if resp.StatusCode != 0 {
		return int64(0), errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, nil
}

func UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	if userClient == nil {
		return nil, errno.ServiceErr
	}
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return &user.UserInfoResponse{}, err
	}
	return resp, nil
}

func UserExist(ctx context.Context, userId int64) (bool, error) {
	if userClient == nil {
		return false, errno.ServiceErr
	}
	resp, err := userClient.UserExist(ctx, &user.UserExistRequest{
		UserId: userId,
	})
	if err != nil {
		return false, err
	}
	return resp.Exist, nil
}
