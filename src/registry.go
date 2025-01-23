package src

import (
	"github.com/oktopriima/marvel/pkg/registry"
	"go.uber.org/dig"
)

func NewRegistry() *dig.Container {
	c := dig.New()

	c = registry.NewConfigRegistry(c)
	c = registry.NewServicesRegistry(c)
	c = registry.NewUsecaseRegistry(c)
	c = registry.NewHandlerRegistry(c)

	return c
}
