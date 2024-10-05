package tracer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oktopriima/marvel/core/config"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/transport"
	"log"
	"net/url"
	"strconv"
	"time"
)

const (
	Byte                uint64 = 1
	maxPacketSize              = int(65000 * Byte)
	MiddlewareTraceName        = "func.middleware"
	HandlerTraceName           = "func.handler"
	UsecaseTraceName           = "func.usecase"
	ServiceTraceName           = "func.repository"
	ModulesTraceName           = "func.modules"
)

type Tracer interface {
	Context() context.Context
	Finish(additionalTags ...map[string]interface{})
}

type tracerImpl struct {
	ctx  context.Context
	span *apm.Span
	tags map[string]interface{}
}

func (t *tracerImpl) Context() context.Context {
	return t.ctx
}

func (t *tracerImpl) Finish(tags ...map[string]interface{}) {
	if tags != nil && t.tags == nil {
		t.tags = make(map[string]interface{})
	}

	for _, tag := range tags {
		for k, v := range tag {
			t.tags[k] = v
		}
	}

	for k, v := range t.tags {
		if v != nil && !t.span.IsExitSpan() {
			val := toString(v)
			t.span.Context.SetLabel(k, val)
		}

		t.captureError(v)
	}

	t.span.End()
}

func (t *tracerImpl) captureError(v interface{}) {
	// capture errors
	if v != nil {
		switch val := v.(type) {
		case error:
			if val != nil {
				apm.CaptureError(t.ctx, val).Send()
			}
		}
	}
	return
}

func toString(v interface{}) (s string) {
	switch val := v.(type) {
	case error:
		if val != nil {
			s = val.Error()
		}
	case string:
		s = val
	case int:
		s = strconv.Itoa(val)
	default:
		b, err := json.Marshal(val)
		if err == nil {
			s = string(b)
		}
	}

	if len(s) >= maxPacketSize {
		s = fmt.Sprintf("overflow, size is: %d, max: %d", len(s), maxPacketSize)
	}
	return
}

// StartTrace starting trace child span from parent span
func StartTrace(ctx context.Context, operationName string, spanType string) Tracer {
	parentSpan := apm.SpanFromContext(ctx)
	opts := apm.SpanOptions{
		Start:  time.Now(),
		Parent: parentSpan.TraceContext(),
	}
	span, newCtx := apm.StartSpanOptions(ctx, operationName, spanType, opts)
	return &tracerImpl{
		ctx:  newCtx,
		span: span,
	}
}

func InitNewTracer(cfg config.AppConfig) *apm.Tracer {
	// close default Tracer
	apm.DefaultTracer().Close()
	trace, err := apm.NewTracer(cfg.APM.ServiceName, cfg.APM.Version)
	if err != nil {
		// handle err
		log.Fatalf("error on call tracer. error %v", err)
	}

	trans, err := transport.NewHTTPTransport(transport.HTTPTransportOptions{})
	if err != nil {
		// handle err
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
	//trace.Transport = trans

	return trace
}
