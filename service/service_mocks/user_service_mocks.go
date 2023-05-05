package service_mocks

import (
	"challange10-dts/dto"
	"challange10-dts/pkg/errs"
	"challange10-dts/service"
)

var (
	CreateNewUser func(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login         func(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
)

type userServiceMock struct{}

func NewUserServiceMock() service.UserService {
	return &userServiceMock{}
}

func (u *userServiceMock) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	return CreateNewUser(payload)
}

func (u *userServiceMock) Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	return Login(payload)
}
