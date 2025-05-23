package main

import (
	"context"
	"github.com/oktopriima/marvel/pkg/kafka"
	bootstrap "github.com/oktopriima/marvel/src/bootstrap/kafka"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/handler"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/router"
)

func main() {
	c := bootstrap.Bootstrap()
	if err := c.Provide(kafka.NewConsumerGroup); err != nil {
		panic(err)
	}

	if err := c.Provide(router.NewRouter); err != nil {
		panic(err)
	}

	if err := c.Provide(handler.NewConsumerHandler); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(r handler.Router) {
		r.KafkaProcessor(context.Background())
	}); err != nil {
		panic(err)
	}
}
