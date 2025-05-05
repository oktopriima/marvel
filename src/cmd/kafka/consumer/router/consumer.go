package router

import (
	"context"
	"fmt"
	"github.com/oktopriima/marvel/pkg/kafka"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/dto"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type KafkaProcessor interface {
	Serve(data *dto.MessageDecoder) error
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan bool
}

func KafkaProcessorHandler(oldCtx context.Context, consumer kafka.Consumer, topic []string, groupId string, handler KafkaProcessor) {
	keepRunning := true
	consumers := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(oldCtx)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
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

	<-consumers.ready

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()
}
