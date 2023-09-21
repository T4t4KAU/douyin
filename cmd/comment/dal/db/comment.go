package db

import (
	"context"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return constants.CommentTableName
}

// AddNewComment add a comment
func AddNewComment(ctx context.Context, comment *Comment) error {
	err := dbConn.WithContext(ctx).Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCommentById delete comment by comment id
func DeleteCommentById(ctx context.Context, commentId int64) error {
	if ok, _ := CheckCommentExist(commentId); !ok {
		return errno.CommentIsNotExistErr
	}
	comment := &Comment{}
	err := dbConn.WithContext(ctx).Where("id = ?", commentId).Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}

// CheckCommentExist 检查评论是否存在
func CheckCommentExist(commentId int64) (bool, error) {
	comment := &Comment{}
	err := dbConn.Where("id = ?", commentId).Find(comment).Error
	if err != nil {
		return false, err
	}
	if comment.ID == 0 {
		return false, nil
	}
	return true, nil
}

// GetCommentListByVideoID 获取视频评论列表
func GetCommentListByVideoID(ctx context.Context, videoId int64) ([]*Comment, error) {
	var commentList []*Comment

	err := dbConn.WithContext(ctx).Table(constants.CommentTableName).
		Where("video_id = ?", videoId).Find(&commentList).Error
	if err != nil {
		return commentList, err
	}
	return commentList, nil
}

// GetCommentCountByVideoID 获取视频评论数
func GetCommentCountByVideoID(ctx context.Context, videoId int64) (int64, error) {
	var sum int64
	err := dbConn.Model(&Comment{}).WithContext(ctx).
		Where("video_id = ?", videoId).Count(&sum).Error
	if err != nil {
		return sum, err
	}
	return sum, nil
}
