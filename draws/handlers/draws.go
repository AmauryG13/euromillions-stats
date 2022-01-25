package handlers

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "github.com/amauryg13/ems/draws/proto"
)

type Draws struct{}

func (e *Draws) Get(ctx context.Context, req *pb.GetRequest, rsp *pb.GetResponse) error {
	log.Infof("Received Draws.Get request: %v", req)
	return nil
}
