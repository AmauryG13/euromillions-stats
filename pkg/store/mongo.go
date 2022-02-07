package store

import (
	"context"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	options store.Options
	client  *mongo.Client
}

func (m *MongoStore) Init(opts ...store.Option) error {
	for _, o := range opts {
		o(&m.options)
	}
	return m.configure()
}

func (m *MongoStore) Options() store.Options {
	return m.options
}

func (m *MongoStore) Close() error {
	err := m.client.Disconnect(context.TODO())

	if err != nil {
		logger.Errorf("%s [store] Error during connection closing")
	}

	return err
}

func (m *MongoStore) Read(key string, doc interface{}, opts ...store.ReadOption) ([]*store.Record, error) {
	db := opts.Database
	table := opts.


	collection := m.getCollection()

}

func (m *MongoStore) Delete(key string, opts ...store.DeleteOption) error {

}

func (m *MongoStore) Write(record *store.Record, opts ...store.WriteOption) error {

}

func (m *MongoStore) List(opts ...store.ListOption) ([]string, error) {

}

func (m *MongoStore) String() string {
	return "mongo"
}

func NewStore(opts ...store.Option) store.Store {
	var options store.Options

	return err
}

func (m *mongoStore) Close() error {
	return m.client.Disconnect(m.options.Context)
}

func (m *mongoStore) Init(opts ...store.Option) error {
	for _, o := range opts {
		o(&m.options)
	}
	// reconfigure
	return m.configure()
}

// List all the known records
func (m *mongoStore) List(opts ...store.ListOption) ([]string, error) {
	var options store.ListOptions
	for _, o := range opts {
		o(&options)
	}


	return []string, nil
}

// Read a single key
func (m *mongoStore) Read(key string, opts ...store.ReadOption) ([]*store.Record, error) {
	var options store.ReadOptions
	for _, o := range opts {
		o(&options)
	}

	return []*store.Record, nil
}

// Write records
func (m *mongoStore) Write(r *store.Record, opts ...store.WriteOption) error {
	var options store.WriteOptions
	for _, o := range opts {
		o(&options)
	}

	return nil
}

// Delete records with keys
func (m *mongoStore) Delete(key string, opts ...store.DeleteOption) error {
	var options store.DeleteOptions
	for _, o := range opts {
		o(&options)
	}

	return nil
}

func (m *mongoStore) Options() store.Options {
	return s.options
}

func (m *mongoStore) String() string {
	return "mongo"
}

// NewStore returns a new micro Store backed by sql
func NewStore(opts ...store.Option) store.Store {
	options := store.Options{
		Database: DefaultDatabase,
		Table:    DefaultTable,
	}

	for _, o := range opts {
		o(&options)
	}

	// new store
	s := new(mongoStore)
	// set the options
	s.options = options
	// best-effort configure the store
	if err := s.configure(); err != nil {
		logger.Error("Error configuring store ", err)
	}

	// return store
	return s
}
