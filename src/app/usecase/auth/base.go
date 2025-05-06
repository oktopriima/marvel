package auth

import (
	"context"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/app/repository/contract"
	"github.com/oktopriima/marvel/src/app/usecase/auth/dto"
	"github.com/oktopriima/thor/jwt"
)

type authenticationUsecase struct {
	userRepo      contract.UserContract
	jwtToken      jwt.AccessToken
	kafkaProducer kafka.Producer
}

type AuthenticationUsecase interface {
	EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error)
}

func NewAuthenticationUsecase(userRepository contract.UserContract, jwtToken jwt.AccessToken, kafkaProducer kafka.Producer) AuthenticationUsecase {
	return &authenticationUsecase{userRepo: userRepository, jwtToken: jwtToken, kafkaProducer: kafkaProducer}
}
