package service

import (
	"context"
	"fmt"

	"github.com/amauryg13/ems/draws/pkg/proto"
	drawsStore "github.com/amauryg13/ems/draws/pkg/store"
	"github.com/amauryg13/ems/pkg/store"
	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jinzhu/copier"
)

func (s Service) GetDraw(ctx context.Context, req *proto.GetDrawRequest, rep *proto.GetDrawResponse) error {
	logger.Infof("Service [%s] GetDraw() arg : %d", s.id, req.GetUuid())

	searchDraw := bson.D{{"uuid", req.GetUuid()}}

	searchResult, err := s.Store.Read(
		searchDraw,
		store.ReadFirst(),
	)

	if err != nil {
		logger.Errorf("Service [%s] GetDraw() error searching draw : %s", s.id, err)
		return err
	}

	var foundDraw drawsStore.Draw
	found := searchResult.Data[0]

	fBytes, _ := bson.Marshal(found)
	_ = bson.Unmarshal(fBytes, &foundDraw)

	var protoDraw proto.Draw

	if err := copier.Copy(&protoDraw, foundDraw); err != nil {
		logger.Errorf("Service [%s] GetDraw() error merging draw : %s", s.id, err)
	}

	rep.Draw = &protoDraw
	return nil
}

func (s Service) AddDraw(ctx context.Context, req *proto.AddDrawRequest, rep *proto.AddDrawResponse) (err error) {
	return nil
}

func (s Service) ListDraw(ctx context.Context, req *proto.ListDrawRequest, rep *proto.ListDrawResponse) (err error) {
	searchDraw := bson.D{}

	searchResult, err := s.Store.Read(
		searchDraw,
	)

	if err != nil {
		logger.Errorf("Service [%s] GetDraw() error", s.id, err)
		return err
	}

	fmt.Println(searchResult.Data)

	return nil
}
