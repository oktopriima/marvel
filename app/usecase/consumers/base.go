package consumers

import (
	"context"
	"github.com/IBM/sarama"
	"log"
)

type consumer struct {
}

func (p *consumer) Init(ctx context.Context, message *sarama.ConsumerMessage, session sarama.ConsumerGroupSession) {
	// TODO Update you code here
	log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

	return
}

type ConsumerUsecase interface {
	Init(ctx context.Context, message *sarama.ConsumerMessage, session sarama.ConsumerGroupSession)
}

func NewConsumerUsecase() ConsumerUsecase {
	return &consumer{}
}
