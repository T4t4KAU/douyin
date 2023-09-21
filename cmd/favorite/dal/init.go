package dal

import (
	"douyin/cmd/favorite/dal/cache"
	"douyin/cmd/favorite/dal/db"
	"douyin/cmd/favorite/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.Init()
}
