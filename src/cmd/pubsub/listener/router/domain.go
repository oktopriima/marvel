package router

import "github.com/oktopriima/marvel/src/app/domain/pubsub/example"

type RegisteredDomain struct {
	ex example.EventProcessor
}

func NewRegisteredDomain(ex example.EventProcessor) *RegisteredDomain {
	return &RegisteredDomain{
		ex: ex,
	}
}
