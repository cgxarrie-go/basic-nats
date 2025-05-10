package subscriber

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"emperror.dev/errors"
	"github.com/nats-io/nats.go"

	"github.com/cgxarrie-go/basic-nats/internal/config"
	"github.com/cgxarrie-go/basic-nats/internal/ports"
)

type subscriber struct {
	name string
}

func New(name string) ports.Service {
	return &subscriber{
		name: name,
	}
}
func (s *subscriber) Start() error {
	cfg := config.Get()

	nc, err := nats.Connect(cfg.NATS.Url)
	if err != nil {
		return errors.Wrapf(err, "subscriber %s connecting to NATS", s.name)
	}

	defer nc.Close()
	log.Printf("subscriber %s connected to NATS\n", s.name)

	_, err = nc.Subscribe(
		config.Get().NATS.Subject,
		func(msg *nats.Msg) {
			log.Printf("subscriber %s message received '%s': %s\n", s.name, msg.Subject, string(msg.Data))
		})

	if err != nil {
		return errors.Wrapf(err, "subscriber %s subscribing to NATS subject '%s'", s.name, config.Get().NATS.Subject)
	}

	log.Printf("subscriber %s subsriberd to subject '%s'", s.name, config.Get().NATS.Subject)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Printf("\nsubscriber %s closing subscription and NATS connection...\n", s.name)

	return nil
}
