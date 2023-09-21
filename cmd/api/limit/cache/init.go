package cache

import (
	"douyin/pkg/constants"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdbLIM     *redis.Client
	expireTime = 1 * time.Second
)

func Init() {
	rdbLIM = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
}
