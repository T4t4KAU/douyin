// Code generated by Kitex v0.6.2. DO NOT EDIT.

package commentservice

import (
	"context"
	comment "douyin/kitex_gen/comment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction": kitex.NewMethodInfo(commentActionHandler, newCommentServiceCommentActionArgs, newCommentServiceCommentActionResult, false),
		"CommentList":   kitex.NewMethodInfo(commentListHandler, newCommentServiceCommentListArgs, newCommentServiceCommentListResult, false),
		"CommentCount":  kitex.NewMethodInfo(commentCountHandler, newCommentServiceCommentCountArgs, newCommentServiceCommentCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentActionArgs)
	realResult := result.(*comment.CommentServiceCommentActionResult)
	success, err := handler.(comment.CommentService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentActionArgs() interface{} {
	return comment.NewCommentServiceCommentActionArgs()
}

func newCommentServiceCommentActionResult() interface{} {
	return comment.NewCommentServiceCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentListArgs)
	realResult := result.(*comment.CommentServiceCommentListResult)
	success, err := handler.(comment.CommentService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentListArgs() interface{} {
	return comment.NewCommentServiceCommentListArgs()
}

func newCommentServiceCommentListResult() interface{} {
	return comment.NewCommentServiceCommentListResult()
}

func commentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentCountArgs)
	realResult := result.(*comment.CommentServiceCommentCountResult)
	success, err := handler.(comment.CommentService).CommentCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentCountArgs() interface{} {
	return comment.NewCommentServiceCommentCountArgs()
}

func newCommentServiceCommentCountResult() interface{} {
	return comment.NewCommentServiceCommentCountResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (r *comment.CommentActionResponse, err error) {
	var _args comment.CommentServiceCommentActionArgs
	_args.Req = req
	var _result comment.CommentServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *comment.CommentListRequest) (r *comment.CommentListResponse, err error) {
	var _args comment.CommentServiceCommentListArgs
	_args.Req = req
	var _result comment.CommentServiceCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentCount(ctx context.Context, req *comment.CommentCountRequest) (r *comment.CommentCountResponse, err error) {
	var _args comment.CommentServiceCommentCountArgs
	_args.Req = req
	var _result comment.CommentServiceCommentCountResult
	if err = p.c.Call(ctx, "CommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}