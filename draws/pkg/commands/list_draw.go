package commands

import (
	"context"
	"fmt"

	"github.com/amauryg13/ems/internal/config"
	"github.com/amauryg13/ems/internal/log"
	"github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/urfave/cli/v2"

	drawsSrv "github.com/amauryg13/ems/api/draws/v0"
)

func ListDraw(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "list",
		Usage:    fmt.Sprintf("add %s", cfg.Service.Name),
		Category: "draws",
		Flags: []cli.Flag{
			&cli.Int64Flag{Name: "page-size", Required: false},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			logger := log.NewLogger(
				log.Name(cfg.Service.Name),
				log.Level(cfg.Log.Level),
				log.File(cfg.Log.File),
			)

			logger.Infof("Starting [command] %s %s", cfg.Service.Name, c.Command.Name)

			srvID := cfg.GRPC.Namespace + "." + cfg.Service.Name
			srv := drawsSrv.NewDrawsService(srvID, grpc.NewClient())

			_, err := srv.ListDraw(ctx, &drawsSrv.ListDrawRequest{
				PageLimit: 50,
			})

			if err != nil {
				fmt.Println(fmt.Errorf("could not list draw %w", err))
				return err
			}

			return nil
		},
	}
}
