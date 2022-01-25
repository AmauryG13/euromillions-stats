package server

import (
	"go-micro.dev/v4"
)

func Server() micro.Service {
	service := micro.NewService(
		micro.Name("test"),
	)

	return service
}
