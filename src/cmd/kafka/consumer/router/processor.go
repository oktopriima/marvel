package router

import (
	"context"
	"github.com/oktopriima/marvel/pkg/kafka/constant"
)

func (p *router) KafkaProcessor(ctx context.Context) {
	KafkaProcessorHandler(ctx, p.kafkaConsumer, constant.KafkaTopic, constant.ConsumerGroup, p.handler)
}
