package main

import (
	"fmt"
	"time"

	"github.com/cgxarrie-go/basic-nats/internal/config"
	"github.com/cgxarrie-go/basic-nats/internal/publisher"
	"github.com/cgxarrie-go/basic-nats/internal/subscriber"
)

func main() {

	cfg := config.Config{
		NATS: config.NatsConfig{
			Url:     "nats://localhost:4222",
			Subject: "my.tanned.eggs",
		},
		NumberOfSubscribers: 5,
		TickerDuration:      5 * time.Second,
	}
	config.Set(&cfg)

	pub := publisher.New()

	for i := 1; i <= cfg.NumberOfSubscribers; i++ {
		sub := subscriber.New(fmt.Sprintf("subscriber %02d", i))

		go func() {
			if err := sub.Start(); err != nil {
				panic(err)
			}
		}()

	}

	if err := pub.Start(); err != nil {
		panic(err)
	}
}
