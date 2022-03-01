package v0

import (
	"context"
	"errors"

	"github.com/amauryg13/ems/internal/config"
	"github.com/amauryg13/ems/internal/log"
	mongoStore "github.com/amauryg13/ems/internal/store"
	"github.com/amauryg13/ems/pkg/store"
)

type Service struct {
	id     string
	log    log.Logger
	Config *config.Config
	Store  store.Store
}

func New(opts ...Option) (s *Service, err error) {
	options := newOptions(opts...)

	logger := options.Logger
	cfg := options.Config

	store, err := createStore(cfg, logger)

	if err != nil {
		return nil, errors.New("could not create a new store")
	}

	s = &Service{
		id:     cfg.GRPC.Namespace + "." + cfg.Service.Name,
		log:    logger,
		Config: cfg,
		Store:  store,
	}
	return
}

func createStore(cfg *config.Config, log log.Logger) (store.Store, error) {
	log.Infof("Service [%s] Store creation", cfg.Service.Name)

	store := mongoStore.NewStore(
		store.Nodes(cfg.Store.Nodes[0]),
		store.Database(cfg.Store.Database),
		store.Collection(cfg.Store.Table),
		store.WithContext(context.Background()),
	)

	return store, nil
}
