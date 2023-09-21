package handlers

import (
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendRegisterResponse 发送注册响应信息
func SendRegisterResponse(c *app.RequestContext, resp user.UserRegisterResponse) {
	c.JSON(consts.StatusOK, resp)
}

// SendUserInfoResponse 发送登录响应信息
func SendUserInfoResponse(c *app.RequestContext, resp user.UserInfoResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendRelationActionResponse(c *app.RequestContext, resp relation.RelationActionResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendRelationFollowListResponse(c *app.RequestContext, resp relation.RelationFollowListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendRelationFollowerListResponse(c *app.RequestContext, resp relation.RelationFollowerListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendRelationFriendListResponse(c *app.RequestContext, resp relation.RelationFriendListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendPublishActionResponse(c *app.RequestContext, resp publish.PublishActionResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendMessageActionResponse(c *app.RequestContext, resp message.MessageActionResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendMessageListResponse(c *app.RequestContext, resp message.MessageListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendFavoriteActionResponse(c *app.RequestContext, resp favorite.FavoriteActionResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendFavoriteListResponse(c *app.RequestContext, resp favorite.FavoriteVideoListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendCommentActionResponse(c *app.RequestContext, resp comment.CommentActionResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendCommentListResponse(c *app.RequestContext, resp comment.CommentListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendFavoriteExistResponse(c *app.RequestContext, resp favorite.FavoriteExistResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendPublishListResponse(c *app.RequestContext, resp publish.PublishListResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendFeedActionResponse(c *app.RequestContext, resp publish.FeedActionResponse) {
	c.JSON(consts.StatusOK, resp)
}

func SendLimitResponse(c *app.RequestContext, resp errno.ErrNo) {
	c.JSON(consts.StatusOK, resp)
}
