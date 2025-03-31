package config

import "time"

type Config struct {
	API
}

type API struct {
	Address string        `default:"8080" envconfig:"ADDRESS"`
	Timeout time.Duration `default:"15s" envconfig:"TIMEOUT"`
}
