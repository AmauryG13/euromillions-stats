package commands

import (
	"context"
	"fmt"

	grpc "github.com/amauryg13/ems/draws/pkg/server"
	"github.com/amauryg13/ems/draws/pkg/service"
	mgrpc "github.com/amauryg13/ems/pkg/services/grpc"

	"github.com/amauryg13/ems/pkg/config"
	"github.com/amauryg13/ems/pkg/log"
	"github.com/oklog/run"
	"github.com/urfave/cli/v2"
)

func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start %s server", "draws"),
		Category: "server",
		Action: func(c *cli.Context) error {
			logger := log.NewLogger(
				log.Name(cfg.Service.Name),
				log.Level(cfg.Log.Level),
				log.File(cfg.Log.File),
			)

			gr := run.Group{}
			ctx, cancel := defineContext(cfg)

			defer cancel()

			grpcServer := grpc.Server(
				mgrpc.Name(cfg.Service.Name),
				mgrpc.Config(cfg),
				mgrpc.Context(ctx),
				mgrpc.Handler(service.Service),
			)

			gr.Add(grpcServer.Run, func(_ error) {
				logger.Info().Str("server", "grpc").Msg("shutting down server")
				cancel()
			})

			return gr.Run()
		},
	}
}

// defineContext sets the context for the extension. If there is a context configured it will create a new child from it,
// if not, it will create a root context that can be cancelled.
func defineContext(cfg *config.Config) (context.Context, context.CancelFunc) {
	return func() (context.Context, context.CancelFunc) {
		if cfg.Context == nil {
			return context.WithCancel(context.Background())
		}
		return context.WithCancel(cfg.Context)
	}()
}
