package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"

	handler "github.com/amauryg13/ems/draws/handlers"
	pb "github.com/amauryg13/ems/draws/proto"
)

var (
	service = "draws"
	version = "latest"
)

func main() {
	service := micro.NewService(
		micro.Name(service),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterDrawsHandler(service.Server(), new(handler.Draws))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
