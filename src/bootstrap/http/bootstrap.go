package http

import (
	"github.com/oktopriima/marvel/src/bootstrap"
	"go.uber.org/dig"
)

func NewBootstrap() *dig.Container {
	c := dig.New()

	c = bootstrap.NewConfig(c)
	c = bootstrap.NewDatabase(c)
	c = bootstrap.NewJWT(c)
	c = NewHttp(c)
	c = bootstrap.NewRepository(c)
	c = NewUsecase(c)
	c = NewHandler(c)

	return c
}
