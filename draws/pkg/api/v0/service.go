package api

import (
	"log"

	"go-micro.dev/v4/config"
)

type DrawService struct {
	id     string
	log    log.Logger
	Config *config.Config
}

func New(opts ...Option) (s *Service, err error) {
	options := newOptions(opts...)

	logger := options.Logger
	config := options.Config
}
