package service

import (
	"github.com/amauryg13/ems/internal/config"
	"github.com/amauryg13/ems/internal/log"
)

type Option func(o *Options)

type Options struct {
	Logger log.Logger
	Config *config.Config
}

func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Logger provides a function to set the Logger option.
func Logger(val log.Logger) Option {
	return func(o *Options) {
		o.Logger = val
	}
}

// Config provides a function to set the Config option.
func Config(val *config.Config) Option {
	return func(o *Options) {
		o.Config = val
	}
}
