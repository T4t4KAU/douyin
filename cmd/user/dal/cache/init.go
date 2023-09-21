package cache

import (
	"douyin/pkg/constants"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdbUsers   *redis.Client
	expireTime = 10 * time.Minute
)

func Init() {
	rdbUsers = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
}
