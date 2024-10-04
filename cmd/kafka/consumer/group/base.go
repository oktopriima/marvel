package group

import (
	"context"
	"errors"
	"flag"
	"github.com/IBM/sarama"
	"github.com/oktopriima/marvel/core/config"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

type ConsumerConfig struct {
	brokers  string
	version  string
	group    string
	topics   string
	assignor string
	oldest   bool
	verbose  bool
	marker   string
}

func NewConsumer(cfg config.AppConfig) *ConsumerConfig {
	return &ConsumerConfig{
		brokers:  cfg.Kafka.Brokers,
		version:  cfg.Kafka.Version,
		group:    cfg.Kafka.Group,
		topics:   cfg.Kafka.Topics,
		assignor: cfg.Kafka.Asignor,
		oldest:   true,
		verbose:  false,
		marker:   cfg.Kafka.Marker,
	}
}

func (c *ConsumerConfig) Serve() {
	c.init()

	keepRunning := true
	log.Println("Starting a new Sarama consumer")

	if c.verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(c.version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	cfg := sarama.NewConfig()
	cfg.Version = version

	switch c.assignor {
	case "sticky":
		cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", c.assignor)
	}

	if c.oldest {
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(c.brokers, ","), c.group, cfg)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(c.topics, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

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
		case <-sigusr1:
			c.toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func (c *ConsumerConfig) init() {
	flag.StringVar(&c.version, "version", sarama.DefaultVersion.String(), "Kafka cluster version")
	flag.BoolVar(&c.oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&c.verbose, "verbose", false, "Sarama logging")
	flag.Parse()

	if len(c.brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(c.topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(c.group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func (c *ConsumerConfig) toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
