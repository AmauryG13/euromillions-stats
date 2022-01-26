package service

import (
	"github.com/amauryg13/ems/pkg/config"
	"github.com/amauryg13/ems/pkg/log"
)

type Service struct {
	id     string
	log    log.Logger
	Config *config.Config
}

func New(opts ...Option) (s *Service, err error) {
	options := newOptions(opts...)

	s = &Service{
		id:     options.Config.GRPC.Namespace + "." + options.Config.Service.Name,
		log:    options.Logger,
		Config: options.Config,
	}

	return
}
