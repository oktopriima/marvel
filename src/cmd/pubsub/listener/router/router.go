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
	pubSubConfig *pubsubrouter.Config
	domain       *RegisteredDomain
}

func NewRouter(cfg config.AppConfig, domain *RegisteredDomain) handler.Router {
	return &router{
		cfg:          cfg,
		pubSubRouter: pubsubRouter(cfg),
		pubSubConfig: pubsubConfig(cfg),
		domain:       domain,
	}
}

func (r *router) eventHandle(svc handler.EventProcessor) pubsubrouter.HandlerFunc {
	return svc.Serve
}

func pubsubRouter(cfg config.AppConfig) *pubsubrouter.Router {
	topics := strings.Split(cfg.PubSub.Topic, ",")
	if len(topics) > 0 && cfg.PubSub.Topic != "" {
		return pubsubrouter.NewRouter()
	}
	return nil
}

func pubsubConfig(cfg config.AppConfig) *pubsubrouter.Config {
	return pubsubrouter.NewPubSubConfig(cfg)
}
