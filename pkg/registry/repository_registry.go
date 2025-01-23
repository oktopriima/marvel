package registry

import (
	"github.com/oktopriima/marvel/src/app/repository"
	"go.uber.org/dig"
)

func NewServicesRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(repository.NewUserRepository); err != nil {
		panic(err)
	}

	return container
}
