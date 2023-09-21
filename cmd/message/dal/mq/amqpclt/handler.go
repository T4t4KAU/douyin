package amqpclt

import (
	"context"
	"strconv"
	"strings"

	"douyin/cmd/message/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// MessageActionAdd 添加用户消息记录
func (a *Actor) MessageActionAdd(ctx context.Context, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 取出用户ID
		params := strings.Split(string(d.Body), "&")
		userId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		toUserId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		Content := params[2]

		klog.Infof("message option(%v,%v,%v)", userId, toUserId, Content)
		// 执行SQL，注必须scan，该SQL才能被执行。
		if _, err = db.AddNewMessage(ctx, &db.Message{
			FromUserId: int64(userId),
			ToUserId:   int64(toUserId),
			Content:    Content,
		}); err != nil {
			// 执行出错，打印日志。
			klog.Errorf("add new message to db：(%v)", err)
		}

	}
}
