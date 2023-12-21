package registry

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/marvel/cmd/http/server"
	"github.com/oktopriima/marvel/core/config"
	"github.com/oktopriima/marvel/core/database"
	"go.uber.org/dig"
)

func NewConfigRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func() config.Config {
		return config.NewConfig()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func() *echo.Echo {
		return echo.New()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(server.NewEchoInstance); err != nil {
		panic(err)
	}

	if err = container.Provide(func(cfg config.Config) database.DBInstance {
		return database.NewDatabaseInstance(cfg)
	}); err != nil {
		panic(err)
	}

	return container
}
