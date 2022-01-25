package commands

import (
	"os"

	emsCli "github.com/amauryg13/ems/pkg/cli"
	"github.com/amauryg13/ems/pkg/config"
	"github.com/urfave/cli/v2"
)

func GetCommands(cfg *config.Config) cli.Commands {
	return []*cli.Command{
		Server(cfg),
	}
}

func Execute(cfg *config.Config) error {
	app := emsCli.DefaultApp(&cli.App{
		Name:     "ems-draws",
		Usage:    "Provide draws for ems",
		Commands: GetCommands(cfg),
	})

	return app.Run(os.Args)
}
