package config

import "github.com/amauryg13/ems/pkg/config"

var (
	path = "./settings.json"
)

func DefaultConfig() *config.Config {
	return &config.Config{
		Service: &config.Service{
			Name: "draws",
		},
		GRPC: &config.GRPC{
			Address:   "127.0.0.1:9010",
			Namespace: "fr.amauryg13.ems",
		},
		Log: &config.Log{
			Level: "info",
			File:  "./draws.log",
		},
	}
}
