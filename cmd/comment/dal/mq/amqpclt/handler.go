package amqpclt

import (
	"context"
	"strconv"
	"strings"

	"douyin/cmd/comment/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// CommentActionAdd 添加用户消息记录
func (a *Actor) CommentActionAdd(comments <-chan amqp.Delivery) {
	for d := range comments {
		// 取出用户ID
		params := strings.Split(string(d.Body), "&")
		userId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		videoId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		Content := params[2]

		klog.Infof("comment db option(%v,%v,%v)", userId, videoId, Content)
		if err = db.AddNewComment(context.Background(), &db.Comment{
			UserId:      int64(userId),
			VideoId:     int64(videoId),
			CommentText: Content,
		}); err != nil {
			klog.Errorf("add new message to db：(%v)", err)
		}
	}
}
