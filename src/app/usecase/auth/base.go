package auth

import (
	"context"
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
	"github.com/oktopriima/marvel/src/app/repository/contract"
	"github.com/oktopriima/marvel/src/app/usecase/auth/dto"
	"github.com/oktopriima/thor/jwt"
)

type authenticationUsecase struct {
	userRepo        contract.UserContract
	jwtToken        jwt.AccessToken
	kafkaProducer   kafka.Producer
	cfg             config.AppConfig
	pubsubPublisher pubsubrouter.Publisher
}

type AuthenticationUsecaseContract interface {
	EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error)
}

func NewAuthenticationUsecase(
	userRepository contract.UserContract,
	jwtToken jwt.AccessToken,
	kafkaProducer kafka.Producer,
	cfg config.AppConfig,
	pubsubPublisher pubsubrouter.Publisher) AuthenticationUsecaseContract {
	return &authenticationUsecase{
		userRepo:        userRepository,
		jwtToken:        jwtToken,
		kafkaProducer:   kafkaProducer,
		cfg:             cfg,
		pubsubPublisher: pubsubPublisher,
	}
}
