package registry

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/cmd/http/server"
	"github.com/oktopriima/marvel/core/config"
	"github.com/oktopriima/marvel/core/database"
	"github.com/oktopriima/thor/jwt"
	"go.uber.org/dig"
)

func NewConfigRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func() config.AppConfig {
		return config.NewAppConfig()
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

	if err = container.Provide(func(cfg config.AppConfig) database.DBInstance {
		return database.NewDatabaseInstance(cfg)
	}); err != nil {
		panic(err)
	}

	// provide auth token
	if err = container.Provide(func(cfg config.AppConfig) jwt.AccessToken {
		return jwt.NewAccessToken(jwt.Request{
			SignatureKey: cfg.Jwt.Key,
			Audience:     "",
			Issuer:       cfg.Jwt.Issuer,
		})
	}); err != nil {
		panic(err)
	}

	return container
}
