package group

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/oktopriima/marvel/app/usecase/consumers"
	"log"
)

const (
	Done = "consumed"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready   chan bool
	marker  string
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
			// process message
			consumer.usecase.Init(session.Context(), message, session)
			session.MarkMessage(message, fmt.Sprintf("%s:%s", consumer.marker, Done))

		case <-session.Context().Done():
			return nil
		}
	}
}
