package bootstrap

import (
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/database"
	"go.uber.org/dig"
)

func NewDatabase(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func(cfg config.AppConfig) database.DBInstance {
		return database.NewDatabaseInstance(cfg)
	}); err != nil {
		panic(err)
	}

	return container
}
