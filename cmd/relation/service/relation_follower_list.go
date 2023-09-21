package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	"sync"
)

type RelationFollowerListService struct {
	ctx context.Context
}

func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{
		ctx: ctx,
	}
}

func (s *RelationFollowerListService) FollowerList(req *relation.RelationFollowerListRequest) ([]*common.User, error) {
	var users []*common.User

	followerIds, err := db.GetFollowerIdList(s.ctx, req.UserId)
	if err != nil {
		return users, err
	}

	userChan := make(chan common.User, len(followerIds))
	errChan := make(chan error, len(followerIds))
	doneChan := make(chan struct{})

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case u := <-userChan:
				users = append(users, &u)
			case <-doneChan:
				return
			}
		}
	}()

	for _, id := range followerIds {
		wg.Add(1)
		go func(userId int64) {
			defer wg.Done()
			resp, e := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
				CurrentUserId: req.CurrentUserId,
				UserId:        userId,
			})
			if e != nil {
				errChan <- e
			} else {
				userChan <- *resp.User
			}
		}(id)
	}

	wg.Wait()
	doneChan <- struct{}{}

	select {
	case err = <-errChan:
		return users, err
	default:
	}

	return users, nil
}
