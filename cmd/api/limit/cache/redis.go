package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const (
	followerSuffix = ":follower"
	followSuffix   = ":follow"
)

func add(c *redis.Client, ctx context.Context, k string, v int64) {
	tx := c.TxPipeline()
	tx.SAdd(ctx, k, v)
	tx.Expire(ctx, k, expireTime)
	tx.Exec(ctx)
}

func del(c *redis.Client, ctx context.Context, k string, v int64) {
	tx := c.TxPipeline()
	tx.SRem(ctx, k, v)
	tx.Expire(ctx, k, expireTime)
	tx.Exec(ctx)
}

func check(c *redis.Client, ctx context.Context, k string) bool {
	if i, _ := c.Exists(ctx, k).Result(); i > 0 {
		return true
	}
	return false
}

func incr(c *redis.Client, ctx context.Context, k string) (int64, error) {
	sum, err := c.SCard(ctx, k).Result()
	if err == nil {
		tx := c.TxPipeline()
		tx.Incr(ctx, k)
		tx.Expire(ctx, k, expireTime)
		tx.Exec(ctx)
		return sum + 1, err
	}
	return sum, err
}
