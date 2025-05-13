package server

import (
	"context"
	"github.com/oktopriima/marvel/src/cmd/pubsub/listener/handler"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Runner struct {
	ready chan bool
}

type Server struct {
	router handler.Router
}

func NewServer(router handler.Router) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run(ctx context.Context) {
	keepRunning := true
	runners := Runner{
		ready: make(chan bool),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Println("starting pubsub consumer")
		s.router.PubSubProcessor(ctx)
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
	wg.Wait()
}
