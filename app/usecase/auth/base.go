package auth

import (
	"context"
	"github.com/oktopriima/marvel/app/repository"
	"github.com/oktopriima/marvel/app/usecase/auth/dto"
	"github.com/oktopriima/thor/jwt"
)

type authenticationUsecase struct {
	userRepo repository.UserRepository
	jwtToken jwt.AccessToken
}

type AuthenticationUsecase interface {
	EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error)
}

func NewAuthenticationUsecase(userRepository repository.UserRepository, jwtToken jwt.AccessToken) AuthenticationUsecase {
	return &authenticationUsecase{userRepo: userRepository, jwtToken: jwtToken}
}
