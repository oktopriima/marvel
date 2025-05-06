package http

import (
	"github.com/oktopriima/marvel/src/app/domain/http/auth"
	"github.com/oktopriima/marvel/src/app/domain/http/users"
	"go.uber.org/dig"
)

func NewHandler(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(auth.NewAuthenticationHandler); err != nil {
		panic(err)
	}
	return container
}
