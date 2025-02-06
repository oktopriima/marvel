package http

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/src/cmd/http/server"
	"go.elastic.co/apm/v2"
	"go.uber.org/dig"
)

func NewHttp(container *dig.Container) *dig.Container {
	var err error
	if err = container.Provide(server.NewEchoInstance); err != nil {
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

	return container
}
