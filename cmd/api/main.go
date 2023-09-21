package main

import (
	"context"
	"douyin/cmd/api/auth"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/pkg/errno"
	"douyin/pkg/logrus"
	tracer "douyin/pkg/trace"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Init() {
	rpc.InitRPC()
	auth.Init()
	tracer.Init("api")
	logrus.InitLogger("api_log.txt")
}

func main() {
	Init()

	r := server.New(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxKeepBodySize(1024*1024*1024*1024),
		server.WithMaxRequestBodySize(1024*1024*1024*1024),
	)

	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"status_code": errno.ServiceErrCode,
				"status_msg":  fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))

	r.Use()

	// 注册路由
	router := r.Group("/douyin")

	router.GET("/feed/", handlers.FeedActionHandler)

	router.GET("/user/", auth.MW.MiddlewareFunc(), handlers.UserInfoHandler)
	router.POST("/user/register/", handlers.RegisterHandler)
	router.POST("/user/login/", auth.MW.LoginHandler)

	relationRouter := router.Group("/relation", auth.MW.MiddlewareFunc())
	relationRouter.POST("/action/", handlers.RelationActionHandler)
	relationRouter.GET("/follow/list/", handlers.RelationFollowListHandler)
	relationRouter.GET("/follower/list/", handlers.RelationFollowerListHandler)
	relationRouter.GET("/friend/list/", handlers.RelationFriendListHandler)

	publishRouter := router.Group("/publish", auth.MW.MiddlewareFunc())
	publishRouter.POST("/action/", handlers.PublishActionHandler)
	publishRouter.GET("/list/", handlers.PublishListHandler)

	messageRouter := router.Group("/message", auth.MW.MiddlewareFunc())
	messageRouter.POST("/action/", handlers.MessageActionHandler)
	messageRouter.GET("/chat/", handlers.MessageListHandler)

	favoriteRouter := router.Group("/favorite", auth.MW.MiddlewareFunc())
	favoriteRouter.POST("/action/", handlers.FavoriteActionHandler)
	favoriteRouter.GET("/list/", handlers.FavoriteListHandler)

	commentRouter := router.Group("/comment", auth.MW.MiddlewareFunc())
	commentRouter.POST("/action/", handlers.CommentActionHandler)
	commentRouter.GET("/list/", handlers.CommentListHandler)

	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})

	r.Spin()
}
