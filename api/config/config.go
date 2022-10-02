package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

var (
	App Config
)

type Config struct {
	MapFilePath         string `env:"MAP_FILE_PATH"`
	AlienMaxStepsNumber int    `env:"ALIEN_MAX_STEPS_NUMBER"`
	LogLevel            string `env:"LOG_LEVEL" env-default:"debug"`
}

func (c *Config) Load() {
	if err := env.Parse(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}
