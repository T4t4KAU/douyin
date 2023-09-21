package dal

import (
	"douyin/cmd/user/dal/cache"
	"douyin/cmd/user/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
