package main

import (
	"context"
	"douyin/cmd/user/pkg"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		r := pkg.BuildBaseResp(errno.ParamErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMsg
		return
	}

	uid, err := service.NewUserRegisterService(ctx).UserRegister(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.UserId = uid

	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		r := pkg.BuildBaseResp(errno.ParamErr)
		resp.StatusCode = r.StatusCode
		resp.StatusMsg = r.StatusMsg
		return
	}

	uid, err := service.NewUserLoginService(ctx).UserLogin(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.UserId = uid

	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp = new(user.UserInfoResponse)

	info, err := service.NewUserInfoService(ctx).UserInfo(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.User = info

	return
}

// UserExist implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserExist(ctx context.Context, req *user.UserExistRequest) (resp *user.UserExistResponse, err error) {
	resp = new(user.UserExistResponse)

	exist, err := service.NewUserExistService(ctx).UserExist(req)
	resp.Exist = exist

	return
}
