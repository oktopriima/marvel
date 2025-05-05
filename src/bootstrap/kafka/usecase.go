package kafka

import (
	"github.com/oktopriima/marvel/src/app/usecase/kafka/consumer/example"
	"go.uber.org/dig"
)

func NewUsecase(container *dig.Container) *dig.Container {
	var err error
	if err = container.Provide(example.NewExampleUsecase); err != nil {
		panic(err)
	}

	return container
}
