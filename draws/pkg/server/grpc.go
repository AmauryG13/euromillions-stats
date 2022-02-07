package grpc

import (
	proto "github.com/amauryg13/ems/draws/pkg/proto"
	"github.com/amauryg13/ems/pkg/services/grpc"
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

	if err := proto.RegisterDrawsServiceHandler(service.Server(), handler); err != nil {
		logger.Fatal("Service handler registration failed")
	}

	return service

}
