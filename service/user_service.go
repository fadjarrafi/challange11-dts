package service

import (
	"challange10-dts/dto"
	"challange10-dts/entity"
	"challange10-dts/pkg/errs"
	"challange10-dts/pkg/helpers"
	"challange10-dts/repository/user_repository"
	"fmt"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewUnauthenticatedError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword {
		return nil, errs.NewUnauthenticatedError("invalid email/password")
	}

	fmt.Println("user =>", user)

	response := dto.LoginResponse{
		Result:     "success",
		Message:    "logged in successfullyy",
		StatusCode: http.StatusOK,
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	userEntity := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	err = userEntity.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(userEntity)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		Message:    "registered successfully",
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}
