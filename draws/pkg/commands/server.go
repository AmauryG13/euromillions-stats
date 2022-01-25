package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/server"
)

func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start %s server"),
		Category: "server",
		Action: func(c *cli.Context) error {

			rpcServer := server.Server(
				grpc.Config(cfg),
			)
		},
	}
}
