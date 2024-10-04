package users

import (
	"context"
	"github.com/oktopriima/marvel/app/contract"
	"github.com/oktopriima/marvel/app/usecase/users/dto"
)

type userUsecase struct {
	userRepo contract.UserContract
}

type UserUsecase interface {
	FindByID(ctx context.Context, ID int64) (*dto.UserResponse, error)
}

func NewUserUsecase(userRepository contract.UserContract) UserUsecase {
	return &userUsecase{
		userRepo: userRepository,
	}
}
