package grpc

import (
	proto "github.com/amauryg13/ems/draws/pkg/proto"
	"github.com/amauryg13/ems/pkg/services/grpc"
)

func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)

	service := grpc.NewService(
		grpc.Name(options.Config.Service.Name),
		grpc.Context(options.Context),
		grpc.Address(options.Config.GRPC.Address),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Logger(options.Logger),
		grpc.Flags(options.Flags...),
	)

	if err := proto.RegisterDrawsServiceHandler(service.Server(), options.Handler); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register service handler")
	}

	return service

}
