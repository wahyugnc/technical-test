package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {
}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	_ = godotenv.Load(filenames...)
	return &ConfigImpl{}
}
