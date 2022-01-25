package commands

import (
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/config"
)

func GetCommands(cfg *config.Config) cli.Commands {
	return []*cli.Command{
		Server(cfg),
	}
}
