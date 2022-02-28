package commands

import (
	"context"
	"fmt"

	"github.com/amauryg13/ems/internal/config"
	"github.com/amauryg13/ems/internal/log"
	"github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	drawsSrv "github.com/amauryg13/ems/api/services/draws/v0"
)

func AddDraw(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "add",
		Usage:    fmt.Sprintf("add %s", cfg.Service.Name),
		Category: "draws",
		Flags: []cli.Flag{
			&cli.Int64Flag{Name: "uuid", Required: true},
			&cli.Int64Flag{Name: "number", Required: true},
			&cli.StringFlag{Name: "day", Required: true},
			&cli.Int64Flag{Name: "cycle", Required: true},
			//&cli.TimestampFlag{Name: "forclusion", Layout: "Jan 2, 2006 at 3:04pm (MST)", Required: true},
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

			_, err := srv.AddDraw(ctx, &drawsSrv.AddDrawRequest{
				Draw: &drawsSrv.Draw{
					Uuid:       c.Int64("uuid"),
					Number:     c.Int64("number"),
					Day:        0,
					Cycle:      int32(c.Int64("cycle")),
					Forclusion: timestamppb.Now(),
				},
			})

			if err != nil {
				fmt.Println(fmt.Errorf("could not create a new draw %w", err))
				return err
			}

			return nil
		},
	}
}
