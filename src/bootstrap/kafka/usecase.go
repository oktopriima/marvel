package kafka

import (
	"github.com/oktopriima/marvel/src/app/usecase/example"
	"github.com/oktopriima/marvel/src/app/usecase/users"
	"go.uber.org/dig"
)

func NewUsecase(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(example.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(users.NewUserUsecase); err != nil {
		panic(err)
	}

	return container
}
