package tracer

import (
	"github.com/oktopriima/marvel/core/config"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/transport"
	"log"
	"net/url"
)

const (
	ProcessTraceName    = "func.processor"
	RepositoryTraceName = "func.repository"
)

func InitNewTracer(cfg config.AppConfig) *apm.Tracer {
	apm.DefaultTracer().Close()
	trace, err := apm.NewTracer(cfg.APM.ServiceName, cfg.APM.Version)
	if err != nil {
		log.Fatalf("error on call tracer. error %v", err)
	}

	trans, err := transport.NewHTTPTransport(transport.HTTPTransportOptions{})
	if err != nil {
		log.Fatalf("error on call http transport. error %v", err)
	}

	trans.SetSecretToken(cfg.APM.SecretToken)
	u, err := url.Parse(cfg.APM.Url)
	if err != nil {
		log.Fatalf("error on parse. error %v", err)
	}

	err = trans.SetServerURL(u)
	if err != nil {
		log.Fatalf("error on set server url. error %v", err)
	}

	return trace
}
