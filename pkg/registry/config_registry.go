package registry

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/src/http/server"
	"github.com/oktopriima/marvel/src/kafka/consumer/group"
	"github.com/oktopriima/thor/jwt"
	"go.elastic.co/apm/v2"
	"go.uber.org/dig"
)

func NewConfigRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func() config.AppConfig {
		return config.NewAppConfig()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func(cfg config.AppConfig) *apm.Tracer {
		return tracer.InitNewTracer(cfg)
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func(t *apm.Tracer) *echo.Echo {
		e := echo.New()
		return e
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

	if err = container.Provide(group.NewConsumer); err != nil {
		panic(err)
	}

	return container
}
