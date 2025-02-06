package kafka

import (
	bootstrap2 "github.com/oktopriima/marvel/src/bootstrap"
	"go.uber.org/dig"
)

func Bootstrap() *dig.Container {
	c := dig.New()

	c = bootstrap2.NewConfig(c)
	c = NewKafka(c)
	c = bootstrap2.NewDatabase(c)
	c = bootstrap2.NewRepository(c)
	c = NewUsecase(c)

	return c
}
