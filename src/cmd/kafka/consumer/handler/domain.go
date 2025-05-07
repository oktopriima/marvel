package handler

import (
	"github.com/labstack/gommon/log"
	"github.com/oktopriima/marvel/pkg/kafka/constant"
	exampleHandler "github.com/oktopriima/marvel/src/app/domain/kafka/example"
	usersHandler "github.com/oktopriima/marvel/src/app/domain/kafka/users"
	"github.com/oktopriima/marvel/src/app/usecase/example"
	"github.com/oktopriima/marvel/src/app/usecase/users"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/messages"
)

type consumerHandler struct {
	userUsecase    users.UserUsecaseContract
	exampleUsecase example.UsecaseContract
}

func NewConsumerHandler(
	userUsecase users.UserUsecaseContract,
	exampleUsecase example.UsecaseContract,
) KafkaProcessorHandler {
	return &consumerHandler{
		userUsecase:    userUsecase,
		exampleUsecase: exampleUsecase,
	}
}

func (c *consumerHandler) Serve(data *messages.MessageDecoder) error {
	switch data.Topic {
	case constant.UserSuccessLoginTopic:
		return usersHandler.NewNotifyLoginHandler(c.userUsecase).Serve(data)

	case constant.ExampleTopic:
		return exampleHandler.NewExampleHandler(c.exampleUsecase).Serve(data)
	default:
		log.Printf("unhandled topic: %s", data.Topic)
		return nil
	}
}
