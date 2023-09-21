package amqpclt

import (
	"context"
	"strconv"
	"strings"

	"douyin/cmd/favorite/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// FavoriteActionAdd 添加用户点赞记录
func (a *Actor) FavoriteActionAdd(ctx context.Context, favorites <-chan amqp.Delivery) {
	for d := range favorites {
		params := strings.Split(string(d.Body), "&")
		userId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		videoId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		action := params[2]

		klog.Infof("favorite db option(%v,%v,%v)", userId, videoId, action)

		if action == "1" {
			if exist, _ := db.CheckFavoriteExist(ctx, int64(userId), int64(videoId)); exist {
				klog.Errorf("favorite exist")
			}
			if _, err = db.AddNewFavorite(ctx, &db.Favorites{
				UserId:  int64(userId),
				VideoId: int64(videoId),
			}); err != nil {
				klog.Errorf("add new favorite to db：(%v)", err)
			}
		} else {
			if exist, _ := db.CheckFavoriteExist(ctx, int64(userId), int64(videoId)); !exist {
				klog.Errorf("favorite no exist")
			}
			if _, err = db.DeleteFavorite(ctx, &db.Favorites{
				UserId:  int64(userId),
				VideoId: int64(videoId),
			}); err != nil {
				klog.Errorf("delete new favorite to db：(%v)", err)
			}
		}
	}
}
