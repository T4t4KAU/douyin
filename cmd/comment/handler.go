package main

import (
	"context"
	"douyin/cmd/comment/pkg"
	"douyin/cmd/comment/service"
	"douyin/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	resp = new(comment.CommentActionResponse)

	c, err := service.NewCommentActionService(ctx).CommentAction(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.Comment = c

	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = new(comment.CommentListResponse)

	comments, err := service.NewCommentListService(ctx).CommentList(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.CommentList = comments

	return
}

// CommentCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentCount(ctx context.Context, req *comment.CommentCountRequest) (resp *comment.CommentCountResponse, err error) {
	resp = new(comment.CommentCountResponse)
	resp.Count, err = service.NewCommentCountService(ctx).CommentCount(req)

	return
}
