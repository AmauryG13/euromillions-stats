package config

import (
	"context"
	"fmt"

	mCfg "go-micro.dev/v4/config"
	mCfgFile "go-micro.dev/v4/config/source/file"
)

type Config struct {
	Log *Log

	Context context.Context

	Service *Service
	GRPC    *GRPC
}

func (c Config) Read(path string) *Config {
	var config Config

	if err := mCfg.Load(mCfgFile.NewSource(
		mCfgFile.WithPath(path),
	)); err != nil {
		fmt.Println(err)
		return &config
	}

	if err := mCfg.Get().Scan(&config); err != nil {
		fmt.Println(err)
		return &config
	}

	return &config
}
