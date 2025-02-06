package router

import (
	"context"
)

func (p *router) KafkaProcessor(ctx context.Context) {
	p.kafkaProcessorHandle(ctx, []string{"test"}, "GROUP_ONE", p.exampleUsecase)
}
