package router

import (
	"context"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/app/usecase/kafka/consumer/example"

	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/handler"
)

type router struct {
	kafkaConsumer  kafka.Consumer
	exampleUsecase example.Usecase
}

func NewRouter(
	consumer kafka.Consumer,
	exampleUsecase example.Usecase,
) handler.Router {
	return &router{
		kafkaConsumer:  consumer,
		exampleUsecase: exampleUsecase,
	}
}

func (p *router) kafkaProcessorHandle(ctx context.Context, topic []string, consumerGroup string, mp KafkaProcessor) {
	KafkaProcessorHandler(ctx, p.kafkaConsumer, topic, consumerGroup, mp)
}
