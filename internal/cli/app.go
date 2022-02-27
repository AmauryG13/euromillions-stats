package emsCli

import (
	"github.com/amauryg13/ems/internal/version"
	"github.com/urfave/cli/v2"
)

func DefaultApp(app *cli.App) *cli.App {

	app.Version = version.String
	app.Compiled = version.Compiled()

	app.Authors = []*cli.Author{
		{
			Name:  "Amaury Guillermin",
			Email: "https://github.com/amauryg13",
		},
	}

	return app
}
