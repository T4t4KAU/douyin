package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type RelationExistService struct {
	ctx context.Context
}

// NewRelationExistService 创建用户关系服务
func NewRelationExistService(ctx context.Context) *RelationExistService {
	return &RelationExistService{
		ctx: ctx,
	}
}

func (s *RelationExistService) RelationExist(req *relation.RelationExistRequest) (bool, bool, error) {
	followExist, err := db.CheckRelationFollowExist(s.ctx, req.CurrentUserId, req.UserId)
	followedExist, err := db.CheckRelationFollowedExist(s.ctx, req.CurrentUserId, req.UserId)
	return followExist, followedExist, err
}
