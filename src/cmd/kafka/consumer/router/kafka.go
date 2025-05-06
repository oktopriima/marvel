package router

import (
	"context"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/app/domain/kafka/example"
	"github.com/oktopriima/marvel/src/app/domain/kafka/users"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/handler"
)

type router struct {
	kafkaConsumer            kafka.Consumer
	exampleHandler           example.Handler
	loginNotificationHandler users.LoginNotificationHandler
}

func NewRouter(
	consumer kafka.Consumer,
	exampleHandler example.Handler,
	loginNotificationHandler users.LoginNotificationHandler,
) handler.Router {
	return &router{
		kafkaConsumer:            consumer,
		exampleHandler:           exampleHandler,
		loginNotificationHandler: loginNotificationHandler,
	}
}

func (p *router) kafkaProcessorHandle(ctx context.Context, topic []string, consumerGroup string, mp KafkaProcessor) {
	KafkaProcessorHandler(ctx, p.kafkaConsumer, topic, consumerGroup, mp)
}
