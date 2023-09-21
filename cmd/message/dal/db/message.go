package db

import (
	"context"
	"douyin/pkg/constants"
	"time"
)

type Message struct {
	ID         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id"`
	FromUserId int64     `json:"from_user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Message) TableName() string {
	return constants.MessageTableName
}

// AddNewMessage 添加新信息
func AddNewMessage(ctx context.Context, msg *Message) (bool, error) {
	err := dbConn.WithContext(ctx).Create(msg).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetMessageListByIdPair 基于时间查询消息
func GetMessageListByIdPair(userId1, userId2 int64, preMsgTime time.Time) ([]Message, error) {
	var messages []Message
	err := dbConn.Where("to_user_id = ? AND from_user_id = ? AND created_at > ?",
		userId1, userId2, preMsgTime).Or("to_user_id = ? AND from_user_id = ? AND created_at > ?",
		userId2, userId1, preMsgTime).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// GetAllMessageListByIdPair 获取所有消息
func GetAllMessageListByIdPair(userId1, userId2 int64) ([]Message, error) {
	var messages []Message

	err := dbConn.Where("to_user_id = ? AND from_user_id = ?",
		userId2, userId1).Or("to_user_id = ? AND from_user_id = ?", userId2, userId1).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// GetLatestMessageByIdPair 查询user1和user2的最新消息
func GetLatestMessageByIdPair(userId1, userId2 int64) (*Message, error) {
	var message Message
	err := dbConn.Where("to_user_id = ? AND from_user_id = ?", userId1, userId2).
		Or("to_user_id = ? AND from_user_id = ?", userId2, userId1).Last(&message).Error
	if err == nil {
		return &message, nil
	} else {
		if err.Error() == "record not found" {
			return nil, nil
		} else {
			return nil, err
		}
	}
}
