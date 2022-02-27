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
	Store   *Store
}

func (c Config) Read(path string) error {

	if err := mCfg.Load(mCfgFile.NewSource(
		mCfgFile.WithPath(path),
	)); err != nil {
		fmt.Println(err)
		return err
	}

	if err := mCfg.Get().Scan(&c); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
