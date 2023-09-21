package dal

import (
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/mq"
)

func Init() {
	db.Init()
	mq.Init()
}
