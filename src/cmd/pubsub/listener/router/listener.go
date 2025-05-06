package router

import (
	"context"
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
	"strings"
)

func (r *router) PubSubProcessor(ctx context.Context) {
	topics := strings.Split(r.cfg.PubSub.Topic, ",")
	if len(topics) > 0 {
		pubsubCfg := pubsubrouter.NewPubSubConfig(r.cfg)
		serv := pubsubrouter.NewServer(ctx, pubsubCfg)

		for _, topic := range topics {
			go serv.Subscribe(topic, r.EventRouter())
		}
	}
}
