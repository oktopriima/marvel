package main

import (
	"context"
	bootstrap "github.com/oktopriima/marvel/bootstrap/kafka"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/handler"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/router"
)

func main() {
	c := bootstrap.Bootstrap()
	if err := c.Invoke(kafka.NewConsumerGroup); err != nil {
		panic(err)
	}

	if err := c.Provide(router.NewRouter); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(r handler.Router) {
		r.KafkaProcessor(context.Background())
	}); err != nil {
		panic(err)
	}
}
