package router

import (
	"context"
	"github.com/oktopriima/marvel/pkg/kafka/constant"
)

func (p *router) KafkaProcessor(ctx context.Context) {
	p.kafkaProcessorHandle(ctx, []string{constant.ExampleTopic}, constant.UserLoginConsumerGroup, p.exampleHandler)
	p.kafkaProcessorHandle(ctx, []string{constant.UserSuccessLoginTopic}, constant.UserLoginConsumerGroup, p.loginNotificationHandler)
}
