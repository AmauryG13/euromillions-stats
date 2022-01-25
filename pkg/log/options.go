package log

type Option func(o *Options)

type Options struct {
	Name  string
	Level string
	File  string
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Name:  "ems",
		Level: "info",
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Name provides a function to set the name option.
func Name(val string) Option {
	return func(o *Options) {
		o.Name = val
	}
}

// Level provides a function to set the level option.
func Level(val string) Option {
	return func(o *Options) {
		o.Level = val
	}
}

// File provides a function to set the color option.
func File(val string) Option {
	return func(o *Options) {
		o.File = val
	}
}
