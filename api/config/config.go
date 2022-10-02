// Package config is implemented to path configuration to the API library.
package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

// An App is a global application configuration instance.
var (
	App Config
)

// Represents API configuration type.
type Config struct {
	MapFilePath         string `env:"MAP_FILE_PATH"`
	AlienMaxStepsNumber int    `env:"ALIEN_MAX_STEPS_NUMBER"`
	LogLevel            string `env:"LOG_LEVEL" env-default:"debug"`
}

// Loads configuration from the config file into environment variables
func (c *Config) Load() {
	if err := env.Parse(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}
