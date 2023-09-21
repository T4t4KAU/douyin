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

type RelationFollowListService struct {
	ctx context.Context
}

func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{
		ctx: ctx,
	}
}

func (s *RelationFollowListService) FollowList(req *relation.RelationFollowListRequest) ([]*common.User, error) {
	var users []*common.User

	followIds, err := db.GetFollowIdList(s.ctx, req.UserId)
	if err != nil {
		return users, err
	}

	userChan := make(chan common.User, len(followIds))
	errChan := make(chan error, len(followIds))
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

	for _, id := range followIds {
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
