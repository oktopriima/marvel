package router

import (
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/handler"
)

type router struct {
	kafkaConsumer kafka.Consumer
	handler       handler.KafkaProcessorHandler
}

func NewRouter(
	consumer kafka.Consumer,
	handler handler.KafkaProcessorHandler,
) handler.Router {
	return &router{
		kafkaConsumer: consumer,
		handler:       handler,
	}
}
