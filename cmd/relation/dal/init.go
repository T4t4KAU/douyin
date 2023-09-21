package dal

import (
	"douyin/cmd/relation/dal/cache"
	"douyin/cmd/relation/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
