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

type RelationFriendListService struct {
	ctx context.Context
}

func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{
		ctx: ctx,
	}
}

func (s *RelationFriendListService) FriendList(req *relation.RelationFriendListRequest) ([]*common.User, error) {
	var users []*common.User

	friendIds, err := db.GetFriendIdList(s.ctx, req.UserId)
	if err != nil {
		return users, err
	}

	userChan := make(chan common.User, len(friendIds))
	errChan := make(chan error, len(friendIds))
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

	for _, id := range friendIds {
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
