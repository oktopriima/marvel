package handler

import "context"

type Router interface {
	KafkaProcessor(ctx context.Context)
}
