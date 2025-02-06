package router

import (
	"context"
	"fmt"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/dto"
	"log"
	"time"
)

type KafkaProcessor interface {
	Serve(data *dto.MessageDecoder) error
}

func KafkaProcessorHandler(ctx context.Context, consumer kafka.Consumer, topic []string, groupId string, handler KafkaProcessor) {
	go func() {
		if consumer == nil {
			log.Fatalf("Please attach lib in router handler")
			return
		}
		consumer.Subscribe(&kafka.ConsumerContext{
			Handler: func(md *kafka.MessageDecoder) {
				data := &dto.MessageDecoder{
					Body:      md.Body,
					Key:       md.Key,
					Message:   md.Message,
					Topic:     md.Topic,
					Partition: md.Partition,
					TimeStamp: md.TimeStamp,
					Offset:    md.Offset,
					Context:   ctx,
				}
				err := handler.Serve(data)
				if err != nil {
					lt := time.Since(md.TimeStamp)
					logger := fmt.Sprintf("group consumer  %s \n"+
						"topic : %s\n"+
						"time : %s", groupId, topic, lt.String())
					log.Printf(logger)
				}
			},
			GroupID: groupId,
			Topics:  topic,
			Context: ctx,
		})
		return
	}()
}
