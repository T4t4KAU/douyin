package rpc

// InitRPC 初始化RPC
func InitRPC() {
	initUserRPC()
	initRelationRPC()
	initPublishRPC()
	initMessageRPC()
	initFavoriteRPC()
	initCommentRPC()
}
