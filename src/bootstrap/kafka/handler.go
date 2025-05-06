package kafka

import (
	"github.com/oktopriima/marvel/src/app/domain/kafka/example"
	"github.com/oktopriima/marvel/src/app/domain/kafka/users"
	"go.uber.org/dig"
)

func NewHandler(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(example.NewExampleHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(users.NewNotifyLoginHandler); err != nil {
		panic(err)
	}

	return container
}
