package registry

import (
	"github.com/oktopriima/marvel/app/handler/users"
	"go.uber.org/dig"
)

func NewHandlerRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserHandler); err != nil {
		panic(err)
	}

	return container
}
