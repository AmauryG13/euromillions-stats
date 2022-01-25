package commands

import (
	"context"
	"fmt"

	"github.com/oklog/run"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/config"
)

func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start %s server", "draws"),
		Category: "server",
		Action: func(c *cli.Context) error {
			logger := logging.Configure(cfg.Service.Name, cfg.Log)

			gr := run.Group{}
			ctx, cancel := defineContext(cfg)

			defer cancel()

			grpcServer := grpc.Server(
				grpc.Name(cfg.Service.Name),
				grpc.Config(cfg),
				grpc.Context(ctx),
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
