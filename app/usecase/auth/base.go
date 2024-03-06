package auth

import (
	"context"
	"github.com/oktopriima/marvel/app/repository"
	"github.com/oktopriima/marvel/app/usecase/auth/dto"
)

type authenticationUsecase struct {
	userRepo repository.UserRepository
}

type AuthenticationUsecase interface {
	EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (*dto.LoginResponse, error)
}

func NewAuthenticationUsecase(userRepository repository.UserRepository) AuthenticationUsecase {
	return &authenticationUsecase{userRepo: userRepository}
}
