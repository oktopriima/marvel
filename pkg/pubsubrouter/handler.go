package pubsubrouter

type Handler interface {
	HandleMessage(m *Message) error
}

// HandlerFunc is an adaptor to allow the use of ordinary functions as message Handlers.
type HandlerFunc func(m *Message) error

// HandleMessage satisfies the Handler interface.
func (h HandlerFunc) HandleMessage(m *Message) error {
	return h(m)
}
