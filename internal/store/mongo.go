package store

import (
	"context"

	"github.com/amauryg13/ems/pkg/store"
	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	options store.Options

	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func (m *mongoStore) Init(opts ...store.Option) error {
	for _, o := range opts {
		o(&m.options)
	}
	// reconfigure
	return m.configure()
}

func (m *mongoStore) Options() store.Options {
	return m.options
}

func (m *mongoStore) Create(value interface{}, opts ...store.CreateOption) (*store.Result, error) {
	var options store.CreateOptions
	for _, o := range opts {
		o(&options)
	}

	var result store.Result

	doc, err := bson.Marshal(value)

	if err != nil {
		logger.Error("[Store] Create() Bson marshal error", err)
		return &result, err
	}

	collection := m.prepare(options.Database, options.Collection)

	res, err := m.insertOne(collection, m.options.Context, doc)

	result.IDs = append(result.IDs, res.InsertedID)

	return &result, err
}

func (m *mongoStore) insertOne(collection *mongo.Collection, ctx context.Context, doc []byte) (*mongo.InsertOneResult, error) {

	result, err := collection.InsertOne(ctx, doc)

	if err != nil {
		logger.Error("[store] Error insertOne", err)
		return &mongo.InsertOneResult{}, nil
	}

	return result, nil
}

func (m *mongoStore) Read(filter interface{}, opts ...store.ReadOption) (*store.Result, error) {
	var options store.ReadOptions
	for _, o := range opts {
		o(&options)
	}

	var res store.Result

	ctx := m.options.Context
	collection := m.prepare(options.Database, options.Collection)

	var results []bson.M
	var err error

	if options.Limit == 1 {
		results, err = m.findOne(collection, ctx, filter, options)
	} else {
		results, err = m.find(collection, ctx, filter, options)
	}

	if err != nil {
		logger.Error("[store] Read() error", err)
		return &res, err
	}

	for _, result := range results {
		schemaRes := m.decodeToSchema(result, options.Schema)
		res.Data = append(res.Data, schemaRes)
	}

	return &res, nil
}

func (m *mongoStore) findOne(col *mongo.Collection, ctx context.Context, filter interface{}, opts store.ReadOptions) ([]bson.M, error) {
	options := options.FindOne()

	if len(opts.Fields) != 0 {
		options.SetProjection(opts.Fields)
	}

	var result []bson.M

	err := col.FindOne(ctx, filter, options).Decode(&result)

	if err != nil {
		logger.Errorf("[store] Error findOne", err)
		return result, err
	}

	return result, nil

}

func (m *mongoStore) find(col *mongo.Collection, ctx context.Context, filter interface{}, opts store.ReadOptions) ([]bson.M, error) {
	options := options.Find()

	options.SetLimit(opts.Limit)
	options.SetSkip(opts.Offset)

	if len(opts.Fields) != 0 {
		options.SetProjection(opts.Fields)
	}

	if len(opts.Sort) != 0 {
		options.SetSort(opts.Sort)
	}

	var results []bson.M

	cursor, err := col.Find(ctx, filter, options)

	if err != nil {
		logger.Error("[store] Error getting cursor", err)
		return results, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M

		if err := cursor.Decode(&result); err != nil {
			logger.Error("[store] Error decoding cursor", err)
			return results, err
		}

		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		logger.Errorf("[store] Error parsing cursor", err)
		return results, err
	}

	return results, err

}

func (m *mongoStore) Update(filter interface{}, changes interface{}, opts ...store.UpdateOption) (*store.Result, error) {
	var options store.UpdateOptions
	for _, o := range opts {
		o(&options)
	}
	var result store.Result

	ctx := m.options.Context
	collection := m.prepare(options.Database, options.Collection)

	res, err := m.updateOne(collection, ctx, filter, changes, options)

	if err != nil {
		logger.Error("[store] Update() error", err)
	}

	result.AffectedRows = res.ModifiedCount

	return &result, nil
}

func (m *mongoStore) updateOne(col *mongo.Collection, ctx context.Context, filter interface{}, doc interface{}, opts store.UpdateOptions) (*mongo.UpdateResult, error) {

	result, err := col.UpdateOne(ctx, filter, doc)

	if err != nil {
		logger.Errorf("[store] Error updateOne", err)
		return &mongo.UpdateResult{}, err
	}

	return result, err
}

func (m *mongoStore) Delete(filter interface{}, opts ...store.DeleteOption) (*store.Result, error) {
	var options store.DeleteOptions
	for _, o := range opts {
		o(&options)
	}
	var result store.Result

	ctx := m.options.Context
	collection := m.prepare(options.Database, options.Collection)

	res, err := m.deleteOne(collection, ctx, filter, options)

	if err != nil {
		logger.Error("[store] Delete() error", err)
		return &result, err
	}

	result.AffectedRows = res.DeletedCount

	return &result, nil
}

func (m *mongoStore) deleteOne(col *mongo.Collection, ctx context.Context, filter interface{}, opts store.DeleteOptions) (*mongo.DeleteResult, error) {

	options := options.Delete().SetHint(bson.D{{"_id", 1}})

	result, err := col.DeleteOne(ctx, filter, options)

	if err != nil {
		logger.Error("[store] Error deleteOne", err)
	}

	return result, err
}

func (m *mongoStore) Close() error {
	err := m.client.Disconnect(m.options.Context)

	if err != nil {
		logger.Warn("[store] Error closing mongo connection", err)
		return err
	}

	return nil
}

func (m *mongoStore) String() string {
	return "mongo"
}

// New returns a new store backed by MongoStore
func NewStore(opts ...store.Option) store.Store {
	options := store.Options{}

	for _, o := range opts {
		o(&options)
	}

	s := new(mongoStore)
	// set the options
	s.options = options

	if err := s.configure(); err != nil {
		logger.Error("[store] Error configuring the store server", err)
	}

	// return
	return s
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

	logger.Debugf("[store] uri : %s", uri)

	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		logger.Errorf("[store] Cannot connect to the node %s", uri)
		return err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Error("[store] Cannot ping the server")
		return err
	}

	m.client = client

	if m.options.Database != "" {
		m.db = client.Database(m.options.Database)
	}

	if m.options.Collection != "" {
		m.collection = m.db.Collection(m.options.Collection)
	}

	return nil
}

func (m *mongoStore) prepare(database string, collection string) *mongo.Collection {
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

func (m *mongoStore) decodeToSchema(doc bson.M, schema interface{}) interface{} {
	docBytes, err := bson.Marshal(doc)

	if err != nil {
		logger.Error("[store] decodeToSchema : bson Marshal", err)
		return schema
	}

	err = bson.Unmarshal(docBytes, &schema)

	if err != nil {
		logger.Error("[store] decodeToSchema : bson Unmarshal", err)
	}

	return schema

}
