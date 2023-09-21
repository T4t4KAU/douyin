package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/pkg"
	"douyin/kitex_gen/message"
	"strconv"
)

type MessageListService struct {
	ctx context.Context
}

func NewMessageListService(ctx context.Context) *MessageListService {
	return &MessageListService{
		ctx: ctx,
	}
}

func (s *MessageListService) MessageList(req *message.MessageListRequest) ([]*message.Message, error) {
	messages := make([]*message.Message, 0)
	fromUserId := req.UserId
	toUserId := req.ToUserId
	preMsgTime := req.PreMsgTime

	dbMessages, err := db.GetMessageListByIdPair(fromUserId, toUserId, pkg.MillTimeStampToTime(preMsgTime))
	if err != nil {
		return messages, err
	}

	for _, dbMessage := range dbMessages {
		messages = append(messages, &message.Message{
			Id:         dbMessage.ID,
			ToUserId:   dbMessage.ToUserId,
			FromUserId: dbMessage.FromUserId,
			Content:    dbMessage.Content,
			CreateTime: strconv.FormatInt(dbMessage.CreatedAt.UnixNano()/1000000, 10),
		})
	}
	return messages, nil
}
