package auth

import (
	"context"
	"github.com/oktopriima/marvel/src/app/contract"
	"github.com/oktopriima/marvel/src/app/usecase/auth/dto"
	"github.com/oktopriima/thor/jwt"
)

type authenticationUsecase struct {
	userRepo contract.UserContract
	jwtToken jwt.AccessToken
}

type AuthenticationUsecase interface {
	EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error)
}

func NewAuthenticationUsecase(userRepository contract.UserContract, jwtToken jwt.AccessToken) AuthenticationUsecase {
	return &authenticationUsecase{userRepo: userRepository, jwtToken: jwtToken}
}
