package store

import (
	"fmt"

	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/bson"
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

	var records []*Record
	var err error

	if options.Limit == 1 {
		records, err = m.findOne(q, options)
	} else {
		records, err = m.find(q, options)
	}
	return records, err
}

func (m *MongoStore) findOne(q *Query, opts QueryOptions) ([]*Record, error) {
	ctx := m.options.Context
	collection := m.setup(opts.Database, opts.Collection)

	options := options.FindOne()

	if len(opts.Fields) != 0 {
		options.SetProjection(opts.Fields)
	}

	var result []byte
	err := collection.FindOne(ctx, q.Filter, options).Decode(&result)

	if err != nil {
		logger.Errorf("[store] Error findOne (db : %s, collection : %s) filter : %s", opts.Database, opts.Collection, q.Filter)
	}

	records := make([]*Record, 1)
	records[0] = &Record{
		Value: result,
	}

	return records, err
}

func (m *MongoStore) find(q *Query, opts QueryOptions) ([]*Record, error) {
	ctx := m.options.Context
	collection := m.setup(opts.Database, opts.Collection)

	options := options.Find()

	options.SetLimit(opts.Limit)
	options.SetSkip(opts.Offset)

	if len(opts.Fields) != 0 {
		options.SetProjection(opts.Fields)
	}

	if len(opts.Sort) != 0 {
		options.SetSort(opts.Sort)
	}

	cursor, err := collection.Find(ctx, q.Filter, options)

	if err != nil {
		logger.Errorf("[store] Error getting cursor (db : %s, collection : %s) filter : %s", opts.Database, opts.Collection, q.Filter)
	}

	var records []*Record

	for cursor.Next(ctx) {
		var result []byte

		if err := cursor.Decode(&result); err != nil {
			logger.Errorf("[store] Error decoding cursor (db : %s, collection : %s) filter : %s", opts.Database, opts.Collection, q.Filter)
		}

		records = append(records, &Record{
			Value: result,
		})
	}

	if err := cursor.Err(); err != nil {
		logger.Errorf("[store] Error parsing cursor (db : %s, collection : %s) filter : %s", opts.Database, opts.Collection, q.Filter)
	}

	cursor.Close(ctx)

	return records, err

}

func (m *MongoStore) Update(q []*Query, opts ...QueryOption) ([]*Record, error) {
	var options QueryOptions
	for _, o := range opts {
		o(&options)
	}

	var records []*Record
	var err error

	return records, err
}

func (m *MongoStore) Delete(q *Query, opts ...QueryOption) (int64, error) {
	var options QueryOptions
	for _, o := range opts {
		o(&options)
	}

	var count int64
	var err error

	if len(q.Doc) > 0 && len(q.Filter) == 0 {
		count, err = m.deleteOne(q, options)
	} else if len(q.Filter) > 0 {
		count, err = m.deleteMany(q, options)
	}

	return count, err
}

func (m *MongoStore) deleteOne(q *Query, opts QueryOptions) (int64, error) {
	ctx := m.options.Context
	collection := m.setup(opts.Database, opts.Collection)

	options := options.Delete().SetHint(bson.D{{"_id", 1}})

	result, err := collection.DeleteOne(ctx, q.Doc, options)

	if err != nil {
		logger.Errorf("[store] Error deleteOne (db : %s, collection : %s) doc : %s", opts.Database, opts.Collection, q.Doc)
	}

	return result.DeletedCount, err
}

func (m *MongoStore) deleteMany(q *Query, opts QueryOptions) (int64, error) {
	ctx := m.options.Context
	collection := m.setup(opts.Database, opts.Collection)

	options := options.Delete().SetHint(bson.D{{"_id", 1}})

	results, err := collection.DeleteMany(ctx, q.Filter, options)

	if err != nil {
		logger.Errorf("[store] Error deleteMany (db : %s, collection : %s) filter : %s", opts.Database, opts.Collection, q.Filter)
	}

	return results.DeletedCount, err
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
