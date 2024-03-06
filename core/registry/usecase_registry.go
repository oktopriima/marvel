package registry

import (
	"github.com/oktopriima/marvel/app/usecase/auth"
	"github.com/oktopriima/marvel/app/usecase/users"
	"go.uber.org/dig"
)

func NewUsecaseRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(auth.NewAuthenticationUsecase); err != nil {
		panic(err)
	}

	return container
}
