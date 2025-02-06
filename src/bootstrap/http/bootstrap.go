package http

import (
	bootstrap2 "github.com/oktopriima/marvel/src/bootstrap"
	"go.uber.org/dig"
)

func NewBootstrap() *dig.Container {
	c := dig.New()

	c = bootstrap2.NewConfig(c)
	c = bootstrap2.NewDatabase(c)
	c = bootstrap2.NewJWT(c)
	c = NewHttp(c)
	c = bootstrap2.NewRepository(c)
	c = NewUsecase(c)
	c = NewHandler(c)

	return c
}
