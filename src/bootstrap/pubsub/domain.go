package pubsub

import (
	"github.com/oktopriima/marvel/src/app/domain/pubsub/example"
	"go.uber.org/dig"
)

func NewDomain(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(example.NewHandler); err != nil {
		panic(err)
	}

	return container
}
