package router

import (
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
	"github.com/oktopriima/marvel/src/cmd/pubsub/listener/handler"
	"strings"
)

type router struct {
	cfg          config.AppConfig
	pubSubRouter *pubsubrouter.Router
}

func NewRouter(cfg config.AppConfig) handler.Router {
	return &router{
		cfg:          cfg,
		pubSubRouter: pubsubRouter(cfg),
	}
}

func (r *router) eventHandle(svc handler.EventProcessor) pubsubrouter.HandlerFunc {
	return svc.Serve
}

func pubsubRouter(cfg config.AppConfig) *pubsubrouter.Router {
	subscription := strings.Split(cfg.PubSub.Topic, ",")
	if len(subscription) > 0 && cfg.PubSub.Topic != "" {
		return pubsubrouter.NewRouter()
	}
	return nil
}
