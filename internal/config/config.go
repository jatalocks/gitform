package config

import (
	"github.com/jatalocks/opsilon/internal/engine"
	"github.com/jatalocks/opsilon/internal/logger"
	"github.com/spf13/viper"
)

type Location struct {
	Path string `mapstructure:"path" validate:"nonzero"`
	Type string `mapstructure:"type" validate:"nonzero"`
}

type Action struct {
	Name     string          `mapstructure:"name" validate:"nonzero"`
	Location Location        `mapstructure:"location" validate:"nonzero"`
	Workflow engine.Workflow `mapstructure:"workflow,omitempty"`
}

type ActionFile struct {
	Actions []Action `mapstructure:"workflows" validate:"nonzero"`
}

var C ActionFile

func GetConfig() []Action {
	err2 := viper.Unmarshal(&C)
	logger.HandleErr(err2)
	return C.Actions
}
