package main

import (
	"context"
	"douyin/cmd/message/pkg"
	"douyin/cmd/message/service"
	"douyin/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, req *message.MessageListRequest) (resp *message.MessageListResponse, err error) {
	resp = new(message.MessageListResponse)

	messages, err := service.NewMessageListService(ctx).MessageList(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.MessageList = messages

	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	resp = new(message.MessageActionResponse)

	err = service.NewMessageActionService(ctx).MessageAction(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg

	return
}
