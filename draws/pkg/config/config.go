package config

import "github.com/amauryg13/ems/internal/config"

const ConfigPath string = "/etc/ems/draws.json"

func DefaultConfig() *config.Config {
	return &config.Config{
		Service: &config.Service{
			Name: "draws",
		},
		GRPC: &config.GRPC{
			Address:   "localhost:9001",
			Namespace: "fr.amauryg13.ems",
		},
		Log: &config.Log{
			Level: "info",
		},
		Store: &config.Store{
			Nodes:    []string{"mongodb://ems-user:password@localhost:27017"},
			Database: "ems",
			Table:    "draws",
		},
	}
}
