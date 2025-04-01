package config

import "time"

type Config struct {
	API
	Redis
}

type API struct {
	Address string        `default:"8080" envconfig:"ADDRESS"`
	Timeout time.Duration `default:"15s" envconfig:"TIMEOUT"`
}

type Redis struct {
	Host     string `default:"8080" envconfig:"REDIS_HOST"`
	Port     string `default:"6379" envconfig:"REDIS_PORT"`
	PoolSize int    `default:"10" envconfig:"REDIS_POOL_SIZE"`
}
