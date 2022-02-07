package store

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ()

type mongoStore struct {
	options Options

	client     *mongo.Client
	collection *mongo.Collection
}

func (m *mongoStore) configure() error {
	ctx := m.options.Context

	if len(m.options.Nodes) == 0 {
		logger.Errorf("[store] No node supplied")
	}

	uri := m.options.Nodes[0]

	if len(m.options.Nodes) > 1 {
		logger.Warnf("[store] Multiple nodes (%d) supplied", len(m.options.Nodes))
	}

	mongoOptions := options.Client()
	mongoOptions.ApplyURI(uri)

	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		logger.Errorf("[store] Cannot connect to the node %s", uri)
	}

	m.client = client

	if m.options.Database != "" && m.options.Table != "" {
		collection := client.Database(m.options.Database).Collection(m.options.Table)
		m.collection = collection
	}

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
