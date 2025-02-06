package bootstrap

import (
	"github.com/oktopriima/marvel/pkg/config"
	"go.uber.org/dig"
)

func NewConfig(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func() config.AppConfig {
		return config.NewAppConfig()
	}); err != nil {
		panic(err)
	}

	return container
}
