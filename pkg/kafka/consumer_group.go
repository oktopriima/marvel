package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type consumerGroup struct {
	config     *sarama.Config
	brokers    []string
	autoCommit bool
}

// NewConsumerGroup return consumer message broker
func NewConsumerGroup(cfg *Config) Consumer {
	m := &consumerGroup{}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()

	if cfg.Version == "" {
		cfg.Version = defaultVersion
	}

	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		log.Fatalf("parse kafka version got: %v", cfg.Version)
	}

	if cfg.SASL.Enable {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = cfg.SASL.User
		config.Net.SASL.Password = cfg.SASL.Password
		config.Net.SASL.Version = sarama.SASLHandshakeV0
		config.Net.SASL.Handshake = true
		config.Net.SASL.Mechanism = sarama.SASLMechanism(cfg.SASL.Mechanism)
		config.Net.TLS.Enable = true
	}

	// The TLS configuration to use for secure connections if
	// enabled (defaults to nil).
	if config.Net.TLS.Enable || cfg.TLS.Enable {
		config.Net.TLS.Config = createTlsConfig(cfg.TLS)
	}

	config.Version = version

	config.Consumer.Offsets.Initial = cfg.Consumer.OffsetInitial
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Session.Timeout = time.Duration(cfg.Consumer.SessionTimeoutSecond) * time.Second
	config.Consumer.Group.Heartbeat.Interval = time.Duration(cfg.Consumer.HeartbeatInterval) * time.Millisecond

	if len(strings.Trim(cfg.Consumer.RebalancedStrategy, " ")) == 0 {
		cfg.Consumer.RebalancedStrategy = sarama.RangeBalanceStrategyName
	}

	st, ok := balanceStrategies[cfg.Consumer.RebalancedStrategy]

	if !ok {
		log.Fatalf("unknown balance strategy: %v", cfg.Consumer.RebalancedStrategy)
	}

	if cfg.ChannelBufferSize > 0 {
		config.ChannelBufferSize = cfg.ChannelBufferSize
	}

	config.Consumer.IsolationLevel = sarama.IsolationLevel(cfg.Consumer.IsolationLevel)

	config.Consumer.Group.Rebalance.Strategy = st
	config.ClientID = cfg.ClientID
	m.brokers = cfg.Brokers
	m.config = config
	m.autoCommit = cfg.Consumer.AutoCommit

	return m
}

// Subscribe message
func (k *consumerGroup) Subscribe(ctx *ConsumerContext) {
	if ctx.GroupID == "" {
		ctx.GroupID = os.Getenv("KAFKA_CLIENT_ID")
	}

	client, err := sarama.NewConsumerGroup(k.brokers, ctx.GroupID, k.config)

	if err != nil {
		log.Fatalf("sarama consumer group err: %v", err)
	}

	handler := NewConsumerHandler(ctx.Handler, k.autoCommit)

	// kafka consumer cfg
	nCtx, cancel := context.WithCancel(ctx.Context)

	defer func() {
		_ = client.Close()
	}()
	// subscriber errors
	go func() {
		for err := range client.Errors() {
			log.Fatalf("[consumer] error %s", err.Error())
		}
	}()

	go func() {
		for {
			select {
			case <-nCtx.Done():
				log.Printf("[consumer] stopped consume topics %v", ctx.Topics)
				return
			default:
				err := client.Consume(nCtx, ctx.Topics, handler)
				if err != nil {
					log.Fatalf("[consumer] topic %v error %s", ctx.Topics, err.Error())
				}
			}
		}
	}()
	log.Printf("[consumer] sarama consumer up and running!... group %s, queue %v", ctx.GroupID, ctx.Topics)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	cancel()
	log.Printf("[consumer] cancelled message with offsets")
}
