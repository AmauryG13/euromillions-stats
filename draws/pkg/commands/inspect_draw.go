package commands

import (
	"context"
	"encoding/json"
	"fmt"

	drawsProto "github.com/amauryg13/ems/draws/pkg/proto"
	"github.com/amauryg13/ems/internal/config"
	"github.com/amauryg13/ems/internal/log"
	"github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/urfave/cli/v2"
)

func InspectDraw(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "inspect",
		Usage:    fmt.Sprintf("inspect a %s", cfg.Service.Name),
		Category: "draws",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "uuid", Aliases: []string{"id"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			logger := log.NewLogger(
				log.Name(cfg.Service.Name),
				log.Level(cfg.Log.Level),
				log.File(cfg.Log.File),
			)

			logger.Infof("Starting [command] %s %s", cfg.Service.Name, c.Command.Name)

			drawsSrvID := cfg.GRPC.Namespace + "." + cfg.Service.Name
			drawsSrv := drawsProto.NewDrawsService(drawsSrvID, grpc.NewClient())

			req := drawsProto.GetDrawRequest{
				Uuid: c.String("uuid"),
			}

			res, err := drawsSrv.GetDraw(ctx, &req)
			generateCliOutput(&res, &logger)

			if err != nil {
				fmt.Println(fmt.Errorf("could not create a new draw %w", err))
				return err
			}

			return nil
		},
	}
}

func generateCliOutput(data interface{}, log *log.Logger) {

	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		log.Errorf("Error")
	}
	fmt.Println(string(jsonData))
}
