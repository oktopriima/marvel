package router

import "github.com/oktopriima/marvel/pkg/pubsubrouter"

func (r *router) EventRouter() *pubsubrouter.Router {
	// define all routes here

	return r.pubSubRouter
}
