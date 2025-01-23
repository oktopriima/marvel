package users

import (
	"context"
	"github.com/oktopriima/marvel/src/app/contract"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
)

type userUsecase struct {
	userRepo contract.UserContract
}

type UserUsecase interface {
	FindByID(ctx context.Context, ID int64) (*dto.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
}

func NewUserUsecase(userRepository contract.UserContract) UserUsecase {
	return &userUsecase{
		userRepo: userRepository,
	}
}
