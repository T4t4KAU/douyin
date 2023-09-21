package dal

import (
	"douyin/cmd/publish/dal/db"
	"douyin/cmd/publish/dal/oss"
)

func Init() {
	db.Init()
	oss.Init()
}
