package limit

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/limit/cache"
	"douyin/pkg/constants"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

var rdlim cache.REQ_FRQ

func GetLmtKey(ipaddr string) string {
	return "click_limit_" + ipaddr
}

func IPLimitMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		clientip := c.ClientIP()
		key := GetLmtKey(clientip)
		if rdlim.CheckFrq(ctx, key) {
			// redis中不存在当前key
			rdlim.AddFrq(ctx, key)
		} else {
			cnt, _ := rdlim.IncrFrq(ctx, key)
			if cnt >= int64(constants.LimitsPerSecond) {
				// 操作过于频繁
				handlers.SendLimitResponse(c, errno.ErrIPLimited)
				c.Abort()
				return
			}
		}

		c.Next(ctx)
	}
}
