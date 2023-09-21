package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
)

type UserExistService struct {
	ctx context.Context
}

func NewUserExistService(ctx context.Context) *UserExistService {
	return &UserExistService{
		ctx: ctx,
	}
}

func (s *UserExistService) UserExist(req *user.UserExistRequest) (bool, error) {
	return db.CheckUserExistById(s.ctx, req.UserId)
}
