package pubsub

import (
	"github.com/oktopriima/marvel/src/app/usecase/example"
	"go.uber.org/dig"
)

func NewUsecase(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(example.NewUsecase); err != nil {
		panic(err)
	}

	return container
}
