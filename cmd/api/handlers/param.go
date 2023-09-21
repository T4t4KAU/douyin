package handlers

type UserParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

type UserInfoParam struct {
	CurrentUserId int64
	UserId        int64
}

type RelationActionParam struct {
	ToUserId   int64
	ActionType int32
}

type RelationListParam struct {
	CurrentUserId int64
	UserId        int64
}

type PublishActionParam struct {
	Data   []byte
	UserId int64
	Title  string
}

type PublishListParam struct {
	CurrentUserId int64
	UserId        int64
}

type MessageActionParam struct {
	UserId     int64
	ToUserId   int64
	ActionType int32
	Content    string
}

type MessageListParam struct {
	UserId   int64
	ToUserId int64
}

type FavoriteActionParam struct {
	UserId     int64
	VideoId    int64
	ActionType int32
}

type CommentActionParam struct {
	UserId      int64
	VideoId     int64
	ActionType  int32
	CommentText string
	CommentId   int64
}

type CommentListParam struct {
	UserId  int64
	VideoId int64
}

type PublishInfoParam struct {
	VideoId int64
}

type FavoriteExistParam struct {
	UserId  int64
	VideoId int64
}

type FavoriteListParam struct {
	CurrentUserId int64
	ToUserId      int64
}

type FeedActionParam struct {
	UserId   int64
	LastTime int64
}
