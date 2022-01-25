package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger(opts ...Option) Logger {
	options := newOptions(opts...)

	switch strings.ToLower(options.Level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	var logger zerolog.Logger

	if options.File != "" {
		f, err := os.OpenFile(options.File, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			print(fmt.Sprintf("file could not be opened for writing: %s. error: %v", options.File, err))
			os.Exit(1)
		}
		logger = logger.Output(f)
	} else {
		logger = zerolog.New(os.Stderr)
	}

	logger = logger.With().
		Str("service", options.Name).
		Timestamp().
		Logger()

	return Logger{
		logger,
	}
}
