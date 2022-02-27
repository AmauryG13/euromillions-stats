package commands

import (
	"os"

	emsCli "github.com/amauryg13/ems/internal/cli"
	"github.com/amauryg13/ems/internal/config"
	"github.com/urfave/cli/v2"
)

func GetCommands(cfg *config.Config) cli.Commands {
	return []*cli.Command{
		Server(cfg),
		InspectDraw(cfg),
		AddDraw(cfg),
		ListDraw(cfg),
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
