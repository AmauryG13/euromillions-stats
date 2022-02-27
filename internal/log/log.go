package log

import (
	"strings"

	"go-micro.dev/v4/logger"
)

type Logger struct {
	*logger.Helper
}

func NewLogger(opts ...Option) Logger {
	options := newOptions(opts...)

	var lvl logger.Level

	switch strings.ToLower(options.Level) {
	case "debug":
		lvl = logger.DebugLevel
	default:
		lvl = logger.InfoLevel
	}

	gmlogger := logger.NewLogger(
		logger.WithLevel(lvl),
	)

	gmhelper := logger.NewHelper(gmlogger)

	return Logger{
		gmhelper,
	}
}
