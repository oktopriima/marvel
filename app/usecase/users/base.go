package users

import (
	"context"
	"github.com/oktopriima/marvel/app/repository"
	"github.com/oktopriima/marvel/app/usecase/users/dto"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

type UserUsecase interface {
	FindByID(ctx context.Context, ID int64) (*dto.UserResponse, error)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepository,
	}
}
