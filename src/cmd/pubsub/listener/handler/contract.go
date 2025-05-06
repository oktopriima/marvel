package handler

import (
	"context"
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
)

type Router interface {
	PubSubProcessor(ctx context.Context)
}

type EventProcessor interface {
	Serve(m *pubsubrouter.Message) error
}
