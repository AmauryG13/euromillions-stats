package grpc

import (
	v0 "github.com/amauryg13/ems/api/draws/v0"
	"github.com/amauryg13/ems/internal/services/grpc"
)

func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)
	logger := options.Logger
	handler := options.Handler

	service := grpc.NewService(
		grpc.Name(options.Config.Service.Name),
		grpc.Context(options.Context),
		grpc.Address(options.Config.GRPC.Address),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Logger(logger),
		grpc.Flags(options.Flags...),
	)

	if err := v0.RegisterDrawsServiceHandler(service.Server(), handler); err != nil {
		logger.Fatal("Service handler registration failed")
	}

	return service

}
