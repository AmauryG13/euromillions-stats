package config

import "context"

type Config struct {
	Log *Log

	Context context.Context

	Service *Service
	GRPC    *GRPC
}
