package registry

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/app/modules/middleware"
	"github.com/oktopriima/marvel/cmd/http/server"
	"github.com/oktopriima/marvel/cmd/kafka/consumer/group"
	"github.com/oktopriima/marvel/core/config"
	"github.com/oktopriima/marvel/core/database"
	"github.com/oktopriima/marvel/core/tracer"
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
		e.Use(middleware.ApmEnabler(t))
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

	if err = container.Provide(func(cfg config.AppConfig) *group.ConsumerConfig {
		return group.NewConsumer(cfg)
	}); err != nil {
		panic(err)
	}

	return container
}
