package main

import (
	"os"

	"github.com/amauryg13/ems/draws/pkg/commands"
	"github.com/amauryg13/ems/draws/pkg/config"
)

func main() {
	if err := commands.Execute(config.DefaultConfig()); err != nil {
		os.Exit(1)
	}
}
