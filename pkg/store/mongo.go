package store

import (
	"fmt"

	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	options Options

	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func (m *MongoStore) Init(opts ...Option) error {
	for _, o := range opts {
		o(&m.options)
	}
	// reconfigure
	return m.configure()
}

func (m *MongoStore) Close() error {
	err := m.client.Disconnect(m.options.Context)

	if err != nil {
		logger.Warnf("[] Error closing mongo connection")
	}

	return err
}

func (m *MongoStore) Options() Options {
	return m.options
}

func (m *MongoStore) String() string {
	return "mongo"
}

func (m *MongoStore) Create(q []*Query, opts ...QueryOption) ([]*Record, error) {
	var options QueryOptions
	for _, o := range opts {
		o(&options)
	}

	isMany := false

	if len(q) > 1 {
		isMany = true
	}

	var records []*Record
	var err error

	if isMany {
		records, err = m.insertMany(q, options)
	} else {
		records, err = m.insertOne(q[0], options)
	}

	return records, err
}

func (m *MongoStore) insertOne(q *Query, opts QueryOptions) ([]*Record, error) {
	ctx := m.options.Context
	collection := m.setup(opts.Database, opts.Collection)

	result, err := collection.InsertOne(ctx, q.Doc)

	if err != nil {
		logger.Errorf("[store] Error insertOne (db : %s, collection : %s) doc : %s", opts.Database, opts.Collection, q.Doc)
	}

	records := make([]*Record, 1)
	records[0] = &Record{
		Key: fmt.Sprint(result.InsertedID),
	}

	return records, err
}

func (m *MongoStore) insertMany(q []*Query, opts QueryOptions) ([]*Record, error) {
	ctx := m.options.Context
	collection := m.setup(opts.Database, opts.Collection)

	docs := make([]interface{}, len(q))
	for idx, query := range q {
		docs[idx] = query.Doc
	}

	results, err := collection.InsertMany(ctx, docs)

	if err != nil {
		logger.Errorf("[store] Error insertMany (db : %s, collection : %s)", opts.Database, opts.Collection)
	}

	records := make([]*Record, len(q))
	for idx, id := range results.InsertedIDs {
		records[idx] = &Record{
			Key: fmt.Sprint(id),
		}
	}

	return records, err
}

func (m *MongoStore) Read(q *Query, opts ...QueryOption) ([]*Record, error) {
	var options QueryOptions
	for _, o := range opts {
		o(&options)
	}

	if options.Limit == 1 {

	} else {

	}
	return queries, nil
}

func (m *MongoStore) findOne(q *Query, opts QueryOptions) ([]*Record, error) {

}

func (m *MongoStore) find(q *Query, opts QueryOptions) ([]*Record, error) {

}

func (m *MongoStore) Update(q []*Query, opts ...QueryOption) ([]*Record, error) {
	var options QueryOptions
	for _, o := range opts {
		o(&options)
	}

	return nil
}

func (m *MongoStore) Delete(q []*Query, opts ...QueryOption) ([]*Record, error) {
	var options QueryOptions
	for _, o := range opts {
		o(&options)
	}

	return nil
}

// New returns a new store backed by MongoStore
func New(opts ...Option) Store {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	// new
	s := new(MongoStore)
	// set the options
	s.options = options
	// best-effort configure the
	if err := s.configure(); err != nil {
		logger.Error("Error configuring  ", err)
	}

	// return
	return s
}

func (m *MongoStore) configure() error {
	ctx := m.options.Context

	if len(m.options.Nodes) == 0 {
		logger.Errorf("[] No node supplied")
	}

	uri := m.options.Nodes[0]

	if len(m.options.Nodes) > 1 {
		logger.Warnf("[] Multiple nodes (%d) supplied", len(m.options.Nodes))
	}

	mongoOptions := options.Client()
	mongoOptions.ApplyURI(uri)

	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		logger.Errorf("[] Cannot connect to the node %s", uri)
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
