package config

import (
	"os"
)

type Config struct {
	Key             string
	StoreBaseURL    string
	StorageFilename string
}

func NewConfig() *Config {
	return &Config{
		Key:             os.Getenv("KEY"),
		StoreBaseURL:    os.Getenv("STORE_BASE_URL"),
		StorageFilename: os.Getenv("STORAGE_FILENAME"),
	}
}
