package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

const userSuffix = ":user"

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

func exist(c *redis.Client, ctx context.Context, k string, v int64) bool {
	if ok, _ := c.SIsMember(ctx, k, v).Result(); ok {
		c.Expire(ctx, k, expireTime)
		return true
	}
	return false
}

func count(c *redis.Client, ctx context.Context, k string) (int64, error) {
	sum, err := c.SCard(ctx, k).Result()
	if err == nil {
		c.Expire(ctx, k, expireTime)
		return sum, err
	}
	return sum, err
}

func get(c *redis.Client, ctx context.Context, k string) []int64 {
	vt := make([]int64, 0)

	v, _ := c.SMembers(ctx, k).Result()
	c.Expire(ctx, k, expireTime)
	for _, vs := range v {
		vI64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, vI64)
	}
	return vt
}
