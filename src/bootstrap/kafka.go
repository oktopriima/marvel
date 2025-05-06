package bootstrap

import (
	"github.com/oktopriima/marvel/pkg/kafka"
	"go.uber.org/dig"
)

func NewKafka(container *dig.Container) *dig.Container {
	var err error
	if err = container.Provide(kafka.NewKafkaConfig); err != nil {
		panic(err)
	}

	if err = container.Provide(kafka.NewProducer); err != nil {
		panic(err)
	}

	return container
}
