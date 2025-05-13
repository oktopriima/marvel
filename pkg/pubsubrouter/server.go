package pubsubrouter

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/oktopriima/marvel/pkg/pubsubrouter/cfg"
	"github.com/oktopriima/marvel/pkg/pubsubrouter/session"
	"log"
	"sync/atomic"
)

type Server struct {
	clients   *pubsub.Client
	ctx       context.Context
	subClient *pubsub.Subscription
	router    *Router
}

func NewSession(ctx context.Context, sess session.Contract) *Server {
	cl, err := cfg.NewClient(sess)
	if err != nil {
		panic(err)
	}

	return &Server{
		clients: cl.Client(),
		ctx:     ctx,
	}
}

func (s *Server) Subscribe(topic string, r *Router) *Server {
	s.subClient = s.clients.Subscription(topic)
	s.router = r
	return s
}

func (s *Server) Start() {
	var received int32
	err := s.subClient.Receive(s.ctx, func(ctx context.Context, msg *pubsub.Message) {
		atomic.AddInt32(&received, 1)
		m := Message{}
		m.Data = msg.Data
		m.Attribute = msg.Attributes
		m.Payload = msg
		m.PublishTime = msg.PublishTime
		m.CtlContext = s.ctx
		m.ID = msg.ID
		err := s.router.HandleMessage(&m)
		if err != nil {
			msg.Ack()
			fmt.Println("error", err.Error())
		}
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
