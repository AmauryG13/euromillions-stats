package store

import (
	"fmt"

	"go-micro.dev/v4/store"
	"go-micro.dev/v4/util/log"
	"gopkg.in/mgo.v2"
)

var (
	mongoUrl string = "mongodb://%v:%v"
)

type MongoStore struct {
	options store.Options

	store *mgo.Session
}

func NewMongoStore(opts ...store.Option) *MongoStore {
	options := store.Options{}

	for _, o := range opts {
		o(&options)
	}

	session, err := mgo.Dial(
		fmt.Sprintf(mongoUrl, options.Nodes[0], options.Database),
	)

	if err != nil {
		log.Error("Mongo connextion failed")
	}

	store := &MongoStore{
		options: options,
		store:   session,
	}

	return store
}
