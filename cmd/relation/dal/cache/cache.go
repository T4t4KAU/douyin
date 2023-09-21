package cache

import (
	"context"
	"strconv"
)

type (
	Follows struct{}
)

// AddFollow 添加用户关注
func (f Follows) AddFollow(ctx context.Context, userId, followerId int64) {
	add(rdbFollows, ctx, strconv.FormatInt(followerId, 10)+followSuffix, userId)
}

// AddFollower 添加用户粉丝
func (f Follows) AddFollower(ctx context.Context, userId, followerId int64) {
	add(rdbFollows, ctx, strconv.FormatInt(userId, 10)+followerSuffix, followerId)
}

// DelFollow 删除用户关注
func (f Follows) DelFollow(ctx context.Context, userId, followerId int64) {
	for i := 0; i < 3; i++ {
		err := del(rdbFollows, ctx, strconv.FormatInt(followerId, 10)+followSuffix, userId)
		if err == nil {
			break
		}
	}
}

// DelFollower 删除用户粉丝
func (f Follows) DelFollower(ctx context.Context, userId, followerId int64) {
	for i := 0; i < 3; i++ {
		err := del(rdbFollows, ctx, strconv.FormatInt(userId, 10)+followerSuffix, followerId)
		if err == nil {
			break
		}
	}
}

func (f Follows) CheckFollow(ctx context.Context, followerId int64) bool {
	return check(rdbFollows, ctx, strconv.FormatInt(followerId, 10)+followSuffix)
}

func (f Follows) CheckFollower(ctx context.Context, userId int64) bool {
	return check(rdbFollows, ctx, strconv.FormatInt(userId, 10)+followerSuffix)
}

// ExistFollow 检查用户关注是否存在
func (f Follows) ExistFollow(ctx context.Context, userId, followerId int64) bool {
	return exist(rdbFollows, ctx, strconv.FormatInt(followerId, 10)+followSuffix, userId)
}

// ExistFollower 检查用户粉丝是否存在
func (f Follows) ExistFollower(ctx context.Context, userId, followerId int64) bool {
	return exist(rdbFollows, ctx, strconv.FormatInt(userId, 10)+followerSuffix, followerId)
}

// CountFollow 统计用户关注数量
func (f Follows) CountFollow(ctx context.Context, followerId int64) (int64, error) {
	return count(rdbFollows, ctx, strconv.FormatInt(followerId, 10)+followSuffix)
}

// CountFollower 统计用户粉丝数量
func (f Follows) CountFollower(ctx context.Context, userId int64) (int64, error) {
	return count(rdbFollows, ctx, strconv.FormatInt(userId, 10)+followerSuffix)
}

// GetFollow 获取用户关注列表
func (f Follows) GetFollow(ctx context.Context, followerId int64) []int64 {
	return get(rdbFollows, ctx, strconv.FormatInt(followerId, 10)+followSuffix)
}

// GetFollower 获取用户粉丝列表
func (f Follows) GetFollower(ctx context.Context, userId int64) []int64 {
	return get(rdbFollows, ctx, strconv.FormatInt(userId, 10)+followerSuffix)
}

// GetFriend 获取用户好友列表
func (f Follows) GetFriend(ctx context.Context, id int64) []int64 {
	friends := make([]int64, 0)
	ks1 := strconv.FormatInt(id, 10) + followSuffix
	ks2 := strconv.FormatInt(id, 10) + followerSuffix

	v, _ := rdbFollows.SInter(ctx, ks1, ks2).Result()
	for _, vs := range v {
		vI64, _ := strconv.ParseInt(vs, 10, 64)
		friends = append(friends, vI64)
	}
	return friends
}
