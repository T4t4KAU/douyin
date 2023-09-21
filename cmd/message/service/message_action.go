package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/message/dal/mq"
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"
	"strconv"
	"strings"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{
		ctx: ctx,
	}
}

func (s *MessageActionService) MessageAction(req *message.MessageActionRequest) error {
	if req.Content == "" {
		return nil
	}
	exist, _ := rpc.UserExist(s.ctx, req.UserId)
	if !exist {
		return errno.UserIsNotExistErr
	}

	// 组装成一个消息 为发送至消息队列作准备
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(req.UserId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.ToUserId)))
	sb.WriteString("&")
	sb.WriteString(req.Content)

	//添加到MQ
	err := mq.AddActor.Publish(s.ctx, sb.String())
	if err != nil {
		return err
	}

	return nil
}
