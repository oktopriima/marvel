package router

import (
	"context"
)

func (p *router) KafkaProcessor(ctx context.Context) {
	p.kafkaProcessorHandle(ctx, []string{
		"topic_one",
	},
		"GROUP_ONE",
		p.exampleUsecase,
	)
}
