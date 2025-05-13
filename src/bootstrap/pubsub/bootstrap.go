package pubsub

import (
	"github.com/oktopriima/marvel/src/bootstrap"
	"go.uber.org/dig"
)

func Bootstrap() *dig.Container {
	c := dig.New()

	c = bootstrap.NewConfig(c)
	c = bootstrap.NewKafka(c)
	c = bootstrap.NewDatabase(c)
	c = bootstrap.NewRepository(c)

	c = NewDomain(c)
	c = NewUsecase(c)

	return c
}
