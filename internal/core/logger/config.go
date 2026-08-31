package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type LoggerConfig struct {
	Level  string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (LoggerConfig, error) {
	var conf LoggerConfig

	if err := envconfig.Process("LOGGER", &conf); err != nil {
		return LoggerConfig{}, fmt.Errorf("process error config: %w", err)
	}

	return conf, nil
}

func NewConfigMust() LoggerConfig {
	conf, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get Logger config: %w", err)
		panic(err)
	}

	return conf
}
