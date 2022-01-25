package server

import (
	"strings"

	gmrpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	gmrpcs "github.com/asim/go-micro/plugins/server/grpc/v4"
	"go-micro.dev/v4"
)

type Service struct {
	micro.Service
}

func NewService(opts ...Option) Service {
	sopts := newOptions(opts...)

	mopts := []micro.Option{
		micro.Server(gmrpcs.NewServer()),
		micro.Client(gmrpcc.NewClient()),

		micro.Address(sopts.Address),
		micro.Name(strings.Join([]string{sopts.Namespace, sopts.Name}, ".")),
		micro.Version(sopts.Version),
		micro.Context(sopts.Context),
		micro.Flags(sopts.Flags...),
	}

	return Service{micro.NewService(mopts...)}
}
