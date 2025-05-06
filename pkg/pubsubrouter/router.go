package pubsubrouter

import (
	"github.com/oktopriima/marvel/pkg/pubsubrouter/cfg"
	"sync"
)

type Router struct {
	sync.Mutex
	handlers map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: map[string]Handler{},
	}
}

func (r *Router) Handle(routes string, h Handler) *Router {
	r.Lock()
	defer r.Unlock()

	r.handlers[routes] = h

	return r
}

func (r *Router) HandleMessage(m *Message) error {
	path := m.Payload.Attributes[cfg.MessageAttributeNameRoute]
	h, okRoute := r.handlers[path]
	if okRoute {
		m.Payload.Ack()
		return h.HandleMessage(m)
	}

	return nil
}
