package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/app/helper"
	"github.com/oktopriima/marvel/src/app/usecase/auth/dto"
	"github.com/oktopriima/thor/jwt"
	"strconv"
	"time"
)

func (a *authenticationUsecase) EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error) {
	user, err := a.userRepo.FindByEmail(request.Email, ctx)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(request.Password, user.Password) {
		return nil, fmt.Errorf("password not match")
	}

	token, err := a.jwtToken.GenerateToken(jwt.Params{
		ID: strconv.Itoa(int(user.Id)),
	})
	if err != nil {
		return nil, err
	}

	go func() {
		key, err := json.Marshal(user)
		if err != nil {
			log.Errorf("error while marshalling user: %v", err)
			return
		}
		err = a.kafkaProducer.Publish(ctx, &kafka.MessageContext{
			Value: &kafka.BodyStateful{
				Body:    user,
				Message: "user-success-login",
				Error:   "",
				Source: &kafka.SourceData{
					Service:       "marvel-api",
					ConsumerGroup: "user-login",
				},
			},
			Key:       key,
			LogId:     fmt.Sprintf("user:login:%s:%d", user.TableName(), user.Id),
			Topic:     "user-auth-login-success",
			Partition: 0,
			Offset:    0,
			TimeStamp: time.Now(),
		})

		if err != nil {
			log.Errorf("error while publishing message to kafka: %v", err)
			return
		}

		return
	}()

	return dto.CreateResponse(token), nil
}
