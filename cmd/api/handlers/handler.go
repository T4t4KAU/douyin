package handlers

import (
	"context"
	"douyin/cmd/api/auth"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"strconv"
	"time"
)

// RegisterHandler 用户注册
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam

	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendRegisterResponse(c, user.UserRegisterResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})

		return
	}

	if len(registerVar.UserName) > constants.UserNameMaxLen {
		err := errno.ErrUserNameOverSize
		SendRegisterResponse(c, user.UserRegisterResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		})
		return
	}
	if ok, err := utils.CheckPassword(registerVar.PassWord); !ok {
		SendRegisterResponse(c, user.UserRegisterResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		})
		return
	}
	// 使用注册rpc
	resp, err := rpc.UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})

	if err != nil {
		SendRegisterResponse(c, *resp)
		return
	}
	auth.MW.LoginHandler(ctx, c)
}

func RelationActionHandler(ctx context.Context, c *app.RequestContext) {
	var relationActionVar RelationActionParam

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	uid, err := strconv.Atoi(c.Query("to_user_id"))
	if err != nil {
		SendRelationActionResponse(c, relation.RelationActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}

	relationActionVar.ToUserId = int64(uid)

	action, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		SendRelationActionResponse(c, relation.RelationActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}

	relationActionVar.ActionType = int32(action)

	resp, err := rpc.RelationAction(ctx, &relation.RelationActionRequest{
		CurrentUserId: userId,
		ToUserId:      relationActionVar.ToUserId,
		ActionType:    relationActionVar.ActionType,
	})
	if err != nil {
		SendRelationActionResponse(c, relation.RelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	SendRelationActionResponse(c, relation.RelationActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

func RelationFollowListHandler(ctx context.Context, c *app.RequestContext) {
	var RelationListVar RelationListParam

	claims := jwt.ExtractClaims(ctx, c)
	RelationListVar.CurrentUserId = int64(claims[constants.IdentityKey].(float64))

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendRelationFollowListResponse(c, relation.RelationFollowListResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}
	RelationListVar.UserId = int64(userId)

	resp, err := rpc.RelationFollowList(ctx, &relation.RelationFollowListRequest{
		CurrentUserId: RelationListVar.CurrentUserId,
		UserId:        RelationListVar.UserId,
	})
	if err != nil {
		SendRelationFollowListResponse(c, relation.RelationFollowListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	}

	SendRelationFollowListResponse(c, *resp)
}

func RelationFollowerListHandler(ctx context.Context, c *app.RequestContext) {
	var RelationListVar RelationListParam

	claims := jwt.ExtractClaims(ctx, c)
	RelationListVar.CurrentUserId = int64(claims[constants.IdentityKey].(float64))

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendRelationFollowerListResponse(c, relation.RelationFollowerListResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}
	RelationListVar.UserId = int64(userId)

	resp, err := rpc.RelationFollowerList(ctx, &relation.RelationFollowerListRequest{
		CurrentUserId: RelationListVar.CurrentUserId,
		UserId:        RelationListVar.UserId,
	})
	if err != nil {
		SendRelationFollowerListResponse(c, relation.RelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	}

	SendRelationFollowerListResponse(c, *resp)
}

func RelationFriendListHandler(ctx context.Context, c *app.RequestContext) {
	var RelationListVar RelationListParam

	claims := jwt.ExtractClaims(ctx, c)
	RelationListVar.CurrentUserId = int64(claims[constants.IdentityKey].(float64))

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendRelationFriendListResponse(c, relation.RelationFriendListResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}
	RelationListVar.UserId = int64(userId)

	resp, err := rpc.RelationFriendList(ctx, &relation.RelationFriendListRequest{
		CurrentUserId: RelationListVar.CurrentUserId,
		UserId:        RelationListVar.UserId,
	})
	if err != nil {
		SendRelationFriendListResponse(c, relation.RelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	}

	SendRelationFriendListResponse(c, *resp)
}

func PublishActionHandler(ctx context.Context, c *app.RequestContext) {
	var PublishActionVar PublishActionParam

	claims := jwt.ExtractClaims(ctx, c)

	PublishActionVar.UserId = int64(claims[constants.IdentityKey].(float64))
	PublishActionVar.Title = c.PostForm("title")
	f, err := c.FormFile("data")
	if err != nil {
		SendPublishActionResponse(c, publish.PublishActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}

	PublishActionVar.Data, _ = utils.ReadFileBytes(f)
	resp, err := rpc.PublishAction(ctx, &publish.PublishActionRequest{
		UserId: PublishActionVar.UserId,
		Data:   PublishActionVar.Data,
		Title:  PublishActionVar.Title,
	})
	if err != nil {
		SendPublishActionResponse(c, publish.PublishActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	SendPublishActionResponse(c, publish.PublishActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

func PublishListHandler(ctx context.Context, c *app.RequestContext) {
	var PublishListVar PublishListParam

	claims := jwt.ExtractClaims(ctx, c)

	PublishListVar.CurrentUserId = int64(claims[constants.IdentityKey].(float64))
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendPublishListResponse(c, publish.PublishListResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}

	PublishListVar.UserId = int64(userId)

	resp, err := rpc.PublishList(ctx, &publish.PublishListRequest{
		CurrentUserId: PublishListVar.CurrentUserId,
		UserId:        PublishListVar.UserId,
	})
	if err != nil {
		return
	}
	SendPublishListResponse(c, publish.PublishListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		VideoList:  resp.VideoList,
	})
}

func MessageActionHandler(ctx context.Context, c *app.RequestContext) {
	var MessageActionVar MessageActionParam

	claims := jwt.ExtractClaims(ctx, c)

	MessageActionVar.UserId = int64(claims[constants.IdentityKey].(float64))
	id, err := strconv.Atoi(c.Query("to_user_id"))
	action, err := strconv.Atoi(c.Query("action_type"))
	content := c.Query("content")
	if err != nil {
		SendMessageActionResponse(c, message.MessageActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	MessageActionVar.ToUserId = int64(id)
	MessageActionVar.ActionType = int32(action)
	MessageActionVar.Content = content

	resp, err := rpc.MessageAction(ctx, &message.MessageActionRequest{
		UserId:     MessageActionVar.UserId,
		ToUserId:   MessageActionVar.ToUserId,
		ActionType: MessageActionVar.ActionType,
		Content:    MessageActionVar.Content,
	})
	SendMessageActionResponse(c, message.MessageActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	})
}

func MessageListHandler(ctx context.Context, c *app.RequestContext) {
	var MessageListVar MessageListParam

	claims := jwt.ExtractClaims(ctx, c)
	id, err := strconv.Atoi(c.Query("to_user_id"))
	if err != nil {
		SendMessageListResponse(c, message.MessageListResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	MessageListVar.UserId = int64(claims[constants.IdentityKey].(float64))
	MessageListVar.ToUserId = int64(id)

	resp, err := rpc.MessageList(ctx, &message.MessageListRequest{
		UserId:     MessageListVar.UserId,
		ToUserId:   MessageListVar.ToUserId,
		PreMsgTime: time.Now().UnixNano() / int64(time.Millisecond),
	})
	SendMessageListResponse(c, message.MessageListResponse{
		StatusCode:  resp.StatusCode,
		StatusMsg:   resp.StatusMsg,
		MessageList: resp.MessageList,
	})
}

func FavoriteActionHandler(ctx context.Context, c *app.RequestContext) {
	var FavoriteActionVar FavoriteActionParam

	claims := jwt.ExtractClaims(ctx, c)
	videoId, err := strconv.Atoi(c.Query("video_id"))
	action, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		SendFavoriteActionResponse(c, favorite.FavoriteActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	FavoriteActionVar.UserId = int64(claims[constants.IdentityKey].(float64))
	FavoriteActionVar.VideoId = int64(videoId)
	FavoriteActionVar.ActionType = int32(action)

	resp, _ := rpc.FavoriteAction(ctx, &favorite.FavoriteActionRequest{
		UserId:     FavoriteActionVar.UserId,
		VideoId:    FavoriteActionVar.VideoId,
		ActionType: FavoriteActionVar.ActionType,
	})
	SendFavoriteActionResponse(c, favorite.FavoriteActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	})
}

func FavoriteListHandler(ctx context.Context, c *app.RequestContext) {
	var FavoriteListVar FavoriteListParam

	claims := jwt.ExtractClaims(ctx, c)
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendFavoriteActionResponse(c, favorite.FavoriteActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	FavoriteListVar.CurrentUserId = int64(claims[constants.IdentityKey].(float64))
	FavoriteListVar.ToUserId = int64(userId)

	resp, _ := rpc.FavoriteVideoList(ctx, &favorite.FavoriteVideoListRequest{
		UserId:   FavoriteListVar.CurrentUserId,
		ToUserId: FavoriteListVar.ToUserId,
	})
	SendFavoriteListResponse(c, *resp)
}

func UserInfoHandler(ctx context.Context, c *app.RequestContext) {
	var UserInfoVar UserInfoParam

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendUserInfoResponse(c, user.UserInfoResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
	}

	claims := jwt.ExtractClaims(ctx, c)
	UserInfoVar.CurrentUserId = int64(claims[constants.IdentityKey].(float64))
	if userId == 0 {
		UserInfoVar.UserId = UserInfoVar.CurrentUserId
	} else {
		UserInfoVar.UserId = int64(userId)
	}
	resp, err := rpc.UserInfo(ctx, &user.UserInfoRequest{
		CurrentUserId: UserInfoVar.CurrentUserId,
		UserId:        UserInfoVar.UserId,
	})
	SendUserInfoResponse(c, user.UserInfoResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		User:       resp.User,
	})
}

func CommentActionHandler(ctx context.Context, c *app.RequestContext) {
	var CommentActionVar CommentActionParam

	claims := jwt.ExtractClaims(ctx, c)
	videoId, err := strconv.Atoi(c.Query("video_id"))
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		SendCommentActionResponse(c, comment.CommentActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	if actionType != 1 {
		commentId, err := strconv.Atoi(c.Query("comment_id"))
		if err != nil {
			SendCommentActionResponse(c, comment.CommentActionResponse{
				StatusCode: errno.ParamErrCode,
				StatusMsg:  errno.ParamErrMsg,
			})
			return
		}
		CommentActionVar.CommentId = int64(commentId)
	}

	CommentActionVar.UserId = int64(claims[constants.IdentityKey].(float64))
	CommentActionVar.VideoId = int64(videoId)
	CommentActionVar.ActionType = int32(actionType)
	CommentActionVar.CommentText = c.Query("comment_text")

	sensitiveWords := []string{"bad", "evil", "dangerous"}

	// 敏感词检测
	if utils.SensitiveWordDetection(sensitiveWords, CommentActionVar.CommentText) {
		SendCommentActionResponse(c, comment.CommentActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	resp, err := rpc.CommentAction(ctx, &comment.CommentActionRequest{
		UserId:      CommentActionVar.UserId,
		VideoId:     CommentActionVar.VideoId,
		ActionType:  CommentActionVar.ActionType,
		CommentText: CommentActionVar.CommentText,
		CommentId:   &CommentActionVar.CommentId,
	})
	SendCommentActionResponse(c, *resp)
}

func CommentListHandler(ctx context.Context, c *app.RequestContext) {
	var CommentListVar CommentListParam

	claims := jwt.ExtractClaims(ctx, c)
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		SendFavoriteActionResponse(c, favorite.FavoriteActionResponse{
			StatusCode: errno.ParamErrCode,
			StatusMsg:  errno.ParamErrMsg,
		})
		return
	}

	CommentListVar.UserId = int64(claims[constants.IdentityKey].(float64))
	CommentListVar.VideoId = int64(videoId)

	resp, err := rpc.CommentList(ctx, &comment.CommentListRequest{
		UserId:  CommentListVar.UserId,
		VideoId: CommentListVar.VideoId,
	})

	SendCommentListResponse(c, comment.CommentListResponse{
		StatusCode:  resp.StatusCode,
		StatusMsg:   resp.StatusMsg,
		CommentList: resp.CommentList,
	})
}

func FeedActionHandler(ctx context.Context, c *app.RequestContext) {
	resp, err := rpc.FeedAction(ctx, &publish.FeedActionRequest{})
	if err != nil {
		SendFeedActionResponse(c, publish.FeedActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	}
	SendFeedActionResponse(c, *resp)
}
