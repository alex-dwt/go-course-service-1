package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "MYAPP"

type ENVConfig struct {
	ServerPort     int    `split_words:"true"`
	LoggerDebug    bool   `split_words:"true"`
	CParam         string `envconfig:"SOME_CUSTOM_PARAM"`
	YamlConfigPath string `split_words:"true"`
}

func NewEnvConfig() (ENVConfig, error) {
	var c ENVConfig

	if err := envconfig.Process(envPrefix, &c); err != nil {
		return ENVConfig{}, fmt.Errorf("envconfig read: %w", err)
	}

	return c, nil
}
