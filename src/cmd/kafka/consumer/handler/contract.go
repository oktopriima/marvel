package handler

import (
	"context"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/messages"
)

type Router interface {
	KafkaProcessor(ctx context.Context)
}

type KafkaProcessorHandler interface {
	Serve(data *messages.MessageDecoder) error
}
