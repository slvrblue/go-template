package config

import (
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	App
	Logger
	HTTPConfig
}

type App struct {
	Name string `env:"app_name" envDefault:"go-template"`
}

type HTTPConfig struct {
	ServerAddr   string        `env:"addr" envDefault:"localhost:8080" valid:",required"`
	WriteTimeout time.Duration `env:"write_timeout" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"read_timeout" envDefault:"15s"`
}

type Logger struct {
	// Level defines logging level (error, warning, info, debug, trace)
	Level string `env:"level" envDefault:"debug" valid:"required,in(error|warning|info|debug|trace)"`

	// EnableCaller is used in the ZAP logger.
	EnableCaller bool `env:"enable_caller" valid:",optional"`

	// HumanReadable forces using the text logger formatter instead of json one.
	HumanReadable bool `env:"human_readable" valid:",optional"`
}

func (c *Config) Parse() (err error) {
	if err = env.Parse(&c.HTTPConfig); err != nil {
		return err
	}

	if err = env.Parse(&c.Logger); err != nil {
		return err
	}

	if err = env.Parse(&c.App); err != nil {
		return err
	}

	return
}
