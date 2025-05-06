package server

import (
	"context"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/handler"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Runner struct {
	ready chan bool
}

type ConsumerServer struct {
	r handler.Router
}

func NewConsumerServer(r handler.Router) *ConsumerServer {
	return &ConsumerServer{r: r}
}

func (s *ConsumerServer) Start() {
	oldCtx := context.Background()
	ctx, cancel := context.WithCancel(oldCtx)

	keepRunning := true
	runners := Runner{
		ready: make(chan bool),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s.r.KafkaProcessor(ctx)
		wg.Wait()
	}()

	<-runners.ready

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
