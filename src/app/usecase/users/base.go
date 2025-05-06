package users

import (
	"context"
	"github.com/oktopriima/marvel/src/app/repository/contract"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
)

type userUsecase struct {
	userRepo contract.UserContract
}

type UserUsecaseContract interface {
	FindByID(ctx context.Context, ID int64) (dto.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (dto.UserResponse, error)
	NotifyLogin(ctx context.Context, req *dto.NotifyLoginRequest) (dto.UserResponse, error)
}

func NewUserUsecase(userRepository contract.UserContract) UserUsecaseContract {
	return &userUsecase{
		userRepo: userRepository,
	}
}
