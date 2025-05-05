package kafka

import (
	"github.com/oktopriima/marvel/src/bootstrap"
	"go.uber.org/dig"
)

func Bootstrap() *dig.Container {
	c := dig.New()

	c = bootstrap.NewConfig(c)
	c = NewKafka(c)
	c = bootstrap.NewDatabase(c)
	c = bootstrap.NewRepository(c)
	c = NewUsecase(c)

	return c
}
