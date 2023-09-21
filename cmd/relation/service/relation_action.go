package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
)

const (
	FOLLOW   = 1
	UNFOLLOW = 2
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService 创建用户关系服务
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{
		ctx: ctx,
	}
}

func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	if req.ActionType != FOLLOW && req.ActionType != UNFOLLOW {
		return errno.ParamErr
	}

	if req.ToUserId == req.CurrentUserId {
		return errno.ParamErr
	}

	exist, err := rpc.UserExist(s.ctx, req.ToUserId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.UserIsNotExistErr
	}

	if req.ActionType == FOLLOW {
		exist, err = db.CheckRelationFollowExist(s.ctx, req.CurrentUserId, req.ToUserId)
		if err != nil {
			return err
		}
		if exist {
			return errno.FollowRelationAlreadyExistErr
		}
		ok, err := db.AddNewRelation(s.ctx, &db.Relation{
			UserId:     req.ToUserId,
			FollowerId: req.CurrentUserId,
		})
		if err != nil {
			return err
		}
		if !ok {
			return errno.ServiceErr
		}
	} else {
		exist, err = db.CheckRelationFollowExist(s.ctx, req.CurrentUserId, req.ToUserId)
		if err != nil {
			return err
		}
		if !exist {
			return errno.FavoriteRelationNotExistErr
		}
		ok, err := db.DeleteRelation(s.ctx, &db.Relation{
			UserId:     req.ToUserId,
			FollowerId: req.CurrentUserId,
		})
		if err != nil {
			return err
		}
		if !ok {
			return errno.ServiceErr
		}
	}
	return nil
}
