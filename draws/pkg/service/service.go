package service

import (
	"context"

	"github.com/amauryg13/ems/draws/pkg/proto"
)

type Service struct{}

func (s Service) GetDraw(ctx context.Context, in *proto.GetDrawRequest, out *proto.Draw) error {

	return nil
}
