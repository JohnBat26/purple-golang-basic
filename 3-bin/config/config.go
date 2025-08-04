package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	return &Config{
		Key: os.Getenv("KEY"),
	}
}
