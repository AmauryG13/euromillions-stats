package store

import (
	"github.com/amauryg13/ems/pkg/store/nosql"
	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	options nosql.Options

	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func (m *MongoStore) Init(opts ...nosql.Option) error {
	for _, o := range opts {
		o(&m.options)
	}
	// reconfigure
	return m.configure()
}

func (m *MongoStore) Close() error {
	err := m.client.Disconnect(m.options.Context)

	if err != nil {
		logger.Warnf("[store] Error closing mongo connection")
	}

	return err
}

func (m *MongoStore) Options() nosql.Options {
	return m.options
}

func (m *MongoStore) String() string {
	return "mongo"
}

func (m *MongoStore) Create(r []*nosql.Record, opts ...nosql.WriteOption) error {
	var options nosql.WriteOptions
	for _, o := range opts {
		o(&options)
	}

	ctx := m.options.Context
	collection := m.setup(options.Database, options.Collection)

	Records, isMany := m.prepare(r)

	var err error

	if isMany {
		_, err = collection.InsertMany(ctx, Records)
	} else {
		_, err = collection.InsertOne(ctx, Records[0])
	}

	return err
}

func (m *MongoStore) Read(r []*nosql.Record, opts ...nosql.ReadOption) ([]*nosql.Record, error) {
	records := make([]*nosql.Record, 1)
	return records, nil
}

func (m *MongoStore) Update(r []*nosql.Record, opts ...nosql.WriteOption) error {
	return nil
}

func (m *MongoStore) Delete(r []*nosql.Record, opts ...nosql.DeleteOption) error {
	return nil
}

// NewStore returns a new Store backed by Mongo
func NewStore(opts ...nosql.Option) nosql.Store {
	options := nosql.Options{}

	for _, o := range opts {
		o(&options)
	}

	// new store
	s := new(MongoStore)
	// set the options
	s.options = options
	// best-effort configure the store
	if err := s.configure(); err != nil {
		logger.Error("Error configuring store ", err)
	}

	// return store
	return s
}

func (m *MongoStore) configure() error {
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

	if m.options.Database != "" {
		m.db = client.Database(m.options.Database)
	}

	if m.options.Collection != "" {
		m.collection = m.db.Collection(m.options.Collection)
	}

	return err
}

func (m *MongoStore) setup(database string, collection string) *mongo.Collection {
	var db *mongo.Database
	var coll *mongo.Collection

	if db = m.db; database != "" {
		db = m.client.Database(database)
	}

	if coll = m.collection; collection != "" {
		coll = db.Collection(collection)
	}

	return coll
}

func (m *MongoStore) prepare(records []*nosql.Record) ([]interface{}, bool) {
	var isMany bool

	if isMany = false; len(records) > 1 {
		isMany = true
	}

	Records := make([]interface{}, len(records))

	for idx, record := range records {
		Records[idx] = record.Value
	}

	return Records, isMany

}
