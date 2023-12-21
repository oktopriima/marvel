package registry

import (
	"github.com/oktopriima/marvel/app/services"
	"go.uber.org/dig"
)

func NewServicesRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(services.NewUserServices); err != nil {
		panic(err)
	}

	return container
}
