package router

import (
	"context"
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
	"log"
	"strings"
)

func (r *router) PubSubProcessor(ctx context.Context) {
	if r.pubSubRouter == nil {
		log.Println("no pubsub router has been configured")
		return
	}

	subs := strings.Split(r.cfg.PubSub.Subscription, ",")
	if len(subs) > 0 {
		server := pubsubrouter.NewServer(ctx, r.pubSubConfig)
		for _, sub := range subs {
			go server.Subscribe(sub, r.EventRouter()).Start()
		}
	}
}
