package router

import (
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
)

func (r *router) EventRouter() *pubsubrouter.Router {
	// define all routes here
	r.pubSubRouter.Handle("/auth", r.eventHandle(r.domain.ex))

	return r.pubSubRouter
}
