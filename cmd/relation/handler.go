package main

import (
	"context"
	"douyin/cmd/relation/pkg"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	err = service.NewRelationActionService(ctx).RelationAction(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg

	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp = new(relation.RelationFollowListResponse)

	resp.UserList, err = service.NewRelationFollowListService(ctx).FollowList(req)

	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg

	return
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp = new(relation.RelationFollowerListResponse)

	resp.UserList, err = service.NewRelationFollowerListService(ctx).FollowerList(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg

	return
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	resp = new(relation.RelationFriendListResponse)

	resp.UserList, err = service.NewRelationFriendListService(ctx).FriendList(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg

	return
}

// RelationCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationCount(ctx context.Context, req *relation.RelationCountRequest) (resp *relation.RelationCountResponse, err error) {
	resp = new(relation.RelationCountResponse)

	followCount, followerCount, err := service.NewRelationCountService(ctx).RelationCount(req)
	resp.FollowCount = followCount
	resp.FollowerCount = followerCount
	return
}

// RelationExist implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationExist(ctx context.Context, req *relation.RelationExistRequest) (resp *relation.RelationExistResponse, err error) {
	resp = new(relation.RelationExistResponse)

	followExist, followedExist, err := service.NewRelationExistService(ctx).RelationExist(req)
	resp.FollowExist = followExist
	resp.FollowedExist = followedExist
	return
}
