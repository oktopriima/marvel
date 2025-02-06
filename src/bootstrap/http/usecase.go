package http

import (
	"github.com/oktopriima/marvel/src/app/usecase/auth"
	"github.com/oktopriima/marvel/src/app/usecase/users"
	"go.uber.org/dig"
)

func NewUsecase(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(auth.NewAuthenticationUsecase); err != nil {
		panic(err)
	}

	return container
}
