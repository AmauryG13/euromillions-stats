package grpc

import (
	"context"

	svc "github.com/amauryg13/ems/draws/pkg/service"
	"github.com/amauryg13/ems/internal/config"
	"github.com/amauryg13/ems/internal/log"
	"github.com/urfave/cli/v2"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Name    string
	Logger  log.Logger
	Context context.Context
	Config  *config.Config
	Flags   []cli.Flag
	Handler *svc.Service
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Name provides a name for the service.
func Name(val string) Option {
	return func(o *Options) {
		o.Name = val
	}
}

// Logger provides a function to set the logger option.
func Logger(val log.Logger) Option {
	return func(o *Options) {
		o.Logger = val
	}
}

// Context provides a function to set the context option.
func Context(val context.Context) Option {
	return func(o *Options) {
		o.Context = val
	}
}

// Config provides a function to set the config option.
func Config(val *config.Config) Option {
	return func(o *Options) {
		o.Config = val
	}
}

// Flags provides a function to set the flags option.
func Flags(val []cli.Flag) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, val...)
	}
}

// Handler provides a function to set the handler option.
func Handler(val *svc.Service) Option {
	return func(o *Options) {
		o.Handler = val
	}
}
