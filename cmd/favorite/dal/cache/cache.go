package cache

import (
	"context"
	"strconv"
)

type (
	Favorites struct{}
)

// AddFavorite 添加点赞记录
func (f Favorites) AddFavorite(ctx context.Context, userId, videoId int64) {
	add(rdbFavorites, ctx, strconv.FormatInt(videoId, 10)+favoriteSuffix, userId)
}

// AddFavorited 添加获赞记录
func (f Favorites) AddFavorited(ctx context.Context, userId, videoId int64) {
	add(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+favoritedSuffix, videoId)
}

// DelFavorite 删除点赞记录
func (f Favorites) DelFavorite(ctx context.Context, userId, videoId int64) {
	for i := 0; i < 5; i++ {
		err := del(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+favoriteSuffix, videoId)
		if err == nil {
			break
		}
	}
}

// DelFavorited 删除点赞记录
func (f Favorites) DelFavorited(ctx context.Context, userId, videoId int64) {
	for i := 0; i < 5; i++ {
		err := del(rdbFavorites, ctx, strconv.FormatInt(videoId, 10)+favoritedSuffix, userId)
		if err == nil {
			break
		}
	}
}

// CheckFavorite 检查点赞记录
func (f Favorites) CheckFavorite(ctx context.Context, userId int64) bool {
	return check(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+favoriteSuffix)
}

// CheckFavorited 检查获赞记录
func (f Favorites) CheckFavorited(ctx context.Context, videoId int64) bool {
	return check(rdbFavorites, ctx, strconv.FormatInt(videoId, 10)+favoritedSuffix)
}

// ExistFavorite 检查用户点赞是否存在
func (f Favorites) ExistFavorite(ctx context.Context, userId, videoId int64) bool {
	return exist(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+favoriteSuffix, videoId)
}

// ExistFavorited 检查用户获赞是否存在
func (f Favorites) ExistFavorited(ctx context.Context, userId, videoId int64) bool {
	return exist(rdbFavorites, ctx, strconv.FormatInt(videoId, 10)+favoritedSuffix, userId)
}

// CountFavorite 统计用户点赞数量
func (f Favorites) CountFavorite(ctx context.Context, userId int64) (int64, error) {
	return count(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+favoriteSuffix)
}

// CountFavorited 统计视频获赞数量
func (f Favorites) CountFavorited(ctx context.Context, videoId int64) (int64, error) {
	return count(rdbFavorites, ctx, strconv.FormatInt(videoId, 10)+favoritedSuffix)
}

// GetFavorite 获取用户点赞列表
func (f Favorites) GetFavorite(ctx context.Context, userId int64) []int64 {
	return get(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+favoriteSuffix)
}

// GetFavorited 获取视频获赞列表
func (f Favorites) GetFavorited(ctx context.Context, videoId int64) []int64 {
	return get(rdbFavorites, ctx, strconv.FormatInt(videoId, 10)+favoritedSuffix)
}
