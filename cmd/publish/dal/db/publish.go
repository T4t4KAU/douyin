package db

import (
	"douyin/pkg/constants"
	"time"
)

type Video struct {
	ID          int64     `json:"id"`
	AuthorID    int64     `json:"author_id"`
	PlayURL     string    `json:"play_url"`
	CoverURL    string    `json:"cover_url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string    `json:"title"`
}

func (Video) TableName() string {
	return constants.PublishTableName
}

// CreateVideo 添加新视频
func CreateVideo(video *Video) (int64, error) {
	if err := dbConn.Create(video).Error; err != nil {
		return 0, err
	}
	return video.ID, nil
}

// GetVideoListByLastTime 基于截止视频获取视频列表
func GetVideoListByLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, constants.VideoFeedCount)
	err := dbConn.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(constants.VideoFeedCount).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

// GetVideoByUserID 获取用户所属的视频
func GetVideoByUserID(userId int64) ([]*Video, error) {
	var videos []*Video
	err := dbConn.Where("author_id = ?", userId).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, err
}

// GetVideoListByVideoIDList 通过视频ID获取视频列表
func GetVideoListByVideoIDList(videoIdList []int64) ([]*Video, error) {
	var videoList []*Video
	var err error
	for _, item := range videoIdList {
		var video *Video
		err = dbConn.Where("id = ?", item).Find(&video).Error
		if err != nil {
			return videoList, err
		}
		videoList = append(videoList, video)
	}

	return videoList, err
}

func GetVideoById(videoId int64) (*Video, error) {
	video := Video{}
	err := dbConn.Where("id = ?", videoId).Find(&video).Error
	return &video, err
}

// GetWorkCount 获取一个用户视频发布的数量
func GetWorkCount(userId int64) (int64, error) {
	var count int64
	err := dbConn.Table(constants.PublishTableName).
		Where("author_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CheckVideoExistById 查询视频是否存在
func CheckVideoExistById(videoId int64) (bool, error) {
	var video Video
	if err := dbConn.Where("id = ?", videoId).Find(&video).Error; err != nil {
		return false, err
	}
	if video == (Video{}) {
		return false, nil
	}
	return true, nil
}
