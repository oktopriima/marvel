package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"time"
)

// Producer represents kafka publisher message topic
type Producer interface {
	Publish(ctx context.Context, msg *MessageContext) error
}

// Consumer represents a Sarama consumer consumer interface
type Consumer interface {
	Subscribe(*ConsumerContext)
}

type MessageContext struct {
	Value     *BodyStateful
	Key       []byte
	LogId     interface{}
	Topic     string
	Partition int32
	Offset    int64
	TimeStamp time.Time
	Verbose   bool
}

type BodyStateful struct {
	Body    interface{} `json:"payload" validate:"required" binding:"required"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Source  *SourceData `json:"source,omitempty"`
}

type SourceData struct {
	Service       string `json:"service"`
	ConsumerGroup string `json:"consumer_group,omitempty"`
}

type ConsumerContext struct {
	Handler MessageProcessorFunc
	Topics  []string
	GroupID string
	Context context.Context
}

var balanceStrategies = map[string]sarama.BalanceStrategy{
	sarama.RoundRobinBalanceStrategyName: sarama.BalanceStrategyRoundRobin,
	sarama.RangeBalanceStrategyName:      sarama.BalanceStrategyRange,
	sarama.StickyBalanceStrategyName:     sarama.BalanceStrategySticky,
}

var offsetInitials = map[string]int64{
	"oldest": -2,
	"newest": -1,
}
