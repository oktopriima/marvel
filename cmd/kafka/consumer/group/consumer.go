package group

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/oktopriima/marvel/app/usecase/consumers"
	"github.com/oktopriima/marvel/core/config"
	"log"
)

const (
	Done = "consumed"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready   chan bool
	cfg     config.AppConfig
	usecase consumers.ConsumerUsecase
}

func (consumer *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

			// process message
			consumer.usecase.Init(session.Context(), message, session)
			session.MarkMessage(message, fmt.Sprintf("%s:%s", consumer.cfg.Kafka.Marker, Done))

		case <-session.Context().Done():
			return nil
		}
	}
}
