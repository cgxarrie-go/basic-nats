package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/cgxarrie-go/basic-nats/internal/config"
	"github.com/cgxarrie-go/basic-nats/internal/publisher"
	"github.com/cgxarrie-go/basic-nats/internal/subscriber"

)

func main() {

	cfg, err := config.Load()
	if err != nil {
		err = errors.Wrap(err, "loading config")
		panic(err)
	}

	config.Set(cfg)

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
