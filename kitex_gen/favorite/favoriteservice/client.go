// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "douyin/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest, callOptions ...callopt.Option) (r *favorite.FavoriteActionResponse, err error)
	FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest, callOptions ...callopt.Option) (r *favorite.FavoriteCountResponse, err error)
	FavoriteExist(ctx context.Context, req *favorite.FavoriteExistRequest, callOptions ...callopt.Option) (r *favorite.FavoriteExistResponse, err error)
	FavoriteCountOfVideo(ctx context.Context, req *favorite.FavoriteCountOfVideoRequest, callOptions ...callopt.Option) (r *favorite.FavoriteCountOfVideoResponse, err error)
	FavoriteVideoList(ctx context.Context, req *favorite.FavoriteVideoListRequest, callOptions ...callopt.Option) (r *favorite.FavoriteVideoListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest, callOptions ...callopt.Option) (r *favorite.FavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest, callOptions ...callopt.Option) (r *favorite.FavoriteCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteCount(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteExist(ctx context.Context, req *favorite.FavoriteExistRequest, callOptions ...callopt.Option) (r *favorite.FavoriteExistResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteExist(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteCountOfVideo(ctx context.Context, req *favorite.FavoriteCountOfVideoRequest, callOptions ...callopt.Option) (r *favorite.FavoriteCountOfVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteCountOfVideo(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteVideoList(ctx context.Context, req *favorite.FavoriteVideoListRequest, callOptions ...callopt.Option) (r *favorite.FavoriteVideoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteVideoList(ctx, req)
}
