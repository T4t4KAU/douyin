package dal

import (
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/dal/mq"
)

func Init() {
	db.Init()
	mq.Init()
}
