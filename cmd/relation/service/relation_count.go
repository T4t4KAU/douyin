package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type RelationCountService struct {
	ctx context.Context
}

// NewRelationCountService 创建用户关系服务
func NewRelationCountService(ctx context.Context) *RelationCountService {
	return &RelationCountService{
		ctx: ctx,
	}
}

func (s *RelationCountService) RelationCount(req *relation.RelationCountRequest) (int64, int64, error) {
	followCount, err := db.GetFollowCount(s.ctx, req.UserId)
	followerCount, err := db.GetFollowerCount(s.ctx, req.UserId)
	return followCount, followerCount, err
}
