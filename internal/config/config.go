package config

import (
	"sync"
	"time"
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

// Config
type Config struct {
	NATS                NatsConfig
	NumberOfSubscribers int
	TickerDuration      time.Duration
}

// NatsConfig
type NatsConfig struct {
	Url     string
	Subject string
}
