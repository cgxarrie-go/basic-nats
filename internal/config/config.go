package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var lock = &sync.Mutex{}

var st *Config

func Set(cfg *Config) {
	lock.Lock()
	defer lock.Unlock()
	st = cfg
}

func Get() *Config {
	lock.Lock()
	defer lock.Unlock()
	if st == nil {
		panic("Config not set")
	}
	return st
}

func Load() (*Config, error) {
	tws := os.Getenv("TICKER_WAIT_SECONDS")
	if tws == "" {
		tws = "1"
	}
	twsDur, err := time.ParseDuration(tws)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid ticker value: %s", tws)
	}

	nos := os.Getenv("NUMBER_OF_SUBSCRIBERS")
	if nos == "" {
		nos = "1"
	}
	nosInt, err := strconv.Atoi(nos)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid number of subscribers: %s", nos)
	}

	cfg := &Config{
		NATS: NatsConfig{
			Url:     os.Getenv("NATS_URL"),
			Subject: os.Getenv("NATS_SUBJECT"),
		},
		NumberOfSubscribers: nosInt,
		TickerDuration:      twsDur,
	}

	return cfg, nil
}

// Config
type Config struct {
	NATS                NatsConfig
	NumberOfSubscribers int           "env:NUMBER_OF_SUBSCRIBERS"
	TickerDuration      time.Duration "env:TICKER_WAIT_SECONDS"
}

// NatsConfig
type NatsConfig struct {
	Url     string "env:NATS_URL"
	Subject string "env:NATS_SUBJECT"
}
