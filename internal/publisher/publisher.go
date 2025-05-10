package publisher

import (
	"fmt"
	"log"
	"time"

	"emperror.dev/errors"
	"github.com/nats-io/nats.go"

	"nats-pub/internal/config"
	"nats-pub/internal/ports"
)

type publisher struct{}

func New() ports.Service {
	return &publisher{}
}
func (p *publisher) Start() error {
	cfg := config.Get()

	nc, err := nats.Connect(cfg.NATS.Url)
	if err != nil {
		return errors.Wrap(err, "publisher connecting to NATS")
	}

	defer nc.Close()
	log.Println("publisher connected to NATS")

	ticker := time.NewTicker(cfg.TickerDuration)
	defer ticker.Stop()

	n := 1

	for range ticker.C {
		msg := fmt.Sprintf("Publishing msg %05d: %s ", n, time.Now().Format(time.RFC3339))
		err := nc.Publish(cfg.NATS.Subject, []byte(msg))
		if err != nil {
			log.Printf("Error publishing messages: %v", err)
		} else {
			log.Printf("Msg Published'%s': %s", cfg.NATS.Subject, msg)
		}
		n++
	}

	return nil
}
