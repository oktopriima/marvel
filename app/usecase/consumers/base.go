package consumers

import (
	"context"
	"github.com/IBM/sarama"
)

type consumer struct {
}

func (p *consumer) Init(ctx context.Context, message *sarama.ConsumerMessage, session sarama.ConsumerGroupSession) {
	session.MarkMessage(message, "")

	return
}

type ConsumerUsecase interface {
	Init(ctx context.Context, message *sarama.ConsumerMessage, session sarama.ConsumerGroupSession)
}

func NewPackageValidationUsecase() ConsumerUsecase {
	return &consumer{}
}
