package db

import (
	"context"
	"douyin/cmd/favorite/dal/cache"
	"douyin/pkg/constants"
	"gorm.io/gorm"
	"time"
)

var rdFavorite cache.Favorites

type Favorites struct {
	ID        int64          `json:"id"`
	UserId    int64          `json:"user_id"`
	VideoId   int64          `json:"video_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

// TableName 表名
func (Favorites) TableName() string {
	return constants.FavoriteTableName
}

// AddNewFavorite 添加点赞
func AddNewFavorite(ctx context.Context, favorite *Favorites) (bool, error) {
	err := dbConn.WithContext(ctx).Create(favorite).Error
	if err != nil {
		return false, err
	}
	// add data to redis
	if rdFavorite.CheckFavorited(ctx, favorite.VideoId) {
		rdFavorite.AddFavorited(ctx, favorite.UserId, favorite.VideoId)
	}
	if rdFavorite.CheckFavorite(ctx, favorite.UserId) {
		rdFavorite.AddFavorite(ctx, favorite.UserId, favorite.VideoId)
	}

	return true, nil
}

// DeleteFavorite 删除点赞
func DeleteFavorite(ctx context.Context, favorite *Favorites) (bool, error) {
	err := dbConn.WithContext(ctx).Where("video_id = ? AND user_id = ?",
		favorite.VideoId, favorite.UserId).Delete(favorite).Error
	if err != nil {
		return false, err
	}

	if rdFavorite.CheckFavorited(ctx, favorite.VideoId) {
		rdFavorite.DelFavorited(ctx, favorite.UserId, favorite.VideoId)
	}
	if rdFavorite.CheckFavorite(ctx, favorite.UserId) {
		rdFavorite.DelFavorite(ctx, favorite.UserId, favorite.VideoId)
	}
	return true, nil
}

// CheckFavoriteExist 查询点赞是否存在
func CheckFavoriteExist(ctx context.Context, userId, videoId int64) (bool, error) {
	if rdFavorite.CheckFavorited(ctx, videoId) {
		return rdFavorite.ExistFavorited(ctx, userId, videoId), nil
	}
	if rdFavorite.CheckFavorite(ctx, userId) {
		return rdFavorite.ExistFavorite(ctx, userId, videoId), nil
	}

	var favorite Favorites
	err := dbConn.WithContext(ctx).Table(constants.FavoriteTableName).
		Where("video_id = ? AND user_id = ?", videoId, userId).Find(&favorite).Error
	if err != nil {
		return false, err
	}
	if favorite == (Favorites{}) {
		return false, nil
	}
	return true, nil
}

// GetTotalFavoritedbCountByAuthorID 查询用户总点赞数
func GetTotalFavoritedbCountByAuthorID(ctx context.Context, authorId int64) (int64, error) {
	var sum int64
	err := dbConn.WithContext(ctx).Table(constants.FavoriteTableName).
		Joins("JOIN videos ON favorites.video_id = videos.id").
		Where("videos.author_id = ?", authorId).
		Not("deleted_at IS NOT NULL").Count(&sum).Error
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func getUserFavoriteIdList(userId int64) ([]int64, error) {
	var favoriteActions []Favorites
	err := dbConn.Where("user_id = ?", userId).Find(&favoriteActions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favoriteActions {
		result = append(result, v.VideoId)
	}
	return result, nil
}

// GetUserFavoriteIdList 获取用户点赞视频列表
func GetUserFavoriteIdList(ctx context.Context, userId int64) ([]int64, error) {
	if rdFavorite.CheckFavorite(ctx, userId) {
		return rdFavorite.GetFavorite(ctx, userId), nil
	}
	return getUserFavoriteIdList(userId)
}

// GetUserFavoriteCountById 获取用户点赞数
func GetUserFavoriteCountById(ctx context.Context, userId int64) (int64, error) {
	if rdFavorite.CheckFavorite(ctx, userId) {
		return rdFavorite.CountFavorite(ctx, userId)
	}
	videos, err := getUserFavoriteIdList(userId)
	if err != nil {
		return 0, err
	}

	// 更新缓存
	go func(user int64, videos []int64) {
		for _, video := range videos {
			rdFavorite.AddFavorited(ctx, user, video)
		}
	}(userId, videos)

	return int64(len(videos)), nil
}

func getVideoFavoriterIdList(videoId int64) ([]int64, error) {
	var favoriteActions []Favorites
	err := dbConn.Table(constants.FavoriteTableName).
		Where("video_id = ?", videoId).Find(&favoriteActions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favoriteActions {
		result = append(result, v.UserId)
	}
	return result, nil
}

// GetVideoFavoriterIdList 获取点赞用户列表
func GetVideoFavoriterIdList(ctx context.Context, videoId int64) ([]int64, error) {
	if rdFavorite.CheckFavorited(ctx, videoId) {
		return rdFavorite.GetFavorited(ctx, videoId), nil
	}
	return getVideoFavoriterIdList(videoId)
}

// GetVideoFavoritedCount 获取视频获赞数
func GetVideoFavoritedCount(ctx context.Context, videoId int64) (int64, error) {
	if rdFavorite.CheckFavorited(ctx, videoId) {
		return rdFavorite.CountFavorited(ctx, videoId)
	}

	favorites, err := getVideoFavoriterIdList(videoId)
	if err != nil {
		return 0, err
	}

	// 更新缓存
	go func(users []int64, video int64) {
		for _, u := range users {
			rdFavorite.AddFavorited(ctx, u, video)
		}
	}(favorites, videoId)
	return int64(len(favorites)), err
}
