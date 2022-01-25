package grpc

import (
	"context"
	"log"

	"github.com/urfave/cli/v2"
)

type Option func(o *Options)

type Options struct {
	//
	Namespace string
	Name      string
	Version   string
	//
	Address string
	//
	Context context.Context
	Logger  log.Logger
	//
	Flags []cli.Flag
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Namespace: "ems.api",
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Logger provides a function to set the logger option.
func Logger(l log.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// Namespace provides a function to set the namespace option.
func Namespace(n string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

// Name provides a function to set the name option.
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Version provides a function to set the version option.
func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

// Address provides a function to set the address option.
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// Context provides a function to set the context option.
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// Flags provides a function to set the flags option.
func Flags(flags ...cli.Flag) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}
