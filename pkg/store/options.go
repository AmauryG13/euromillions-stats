package store

import (
	"context"
)

// Options contains configuration for the Store
type Options struct {
	// Nodes contains the addresses or other connection information of the backing storage.
	// For example, an etcd implementation would contain the nodes of the cluster.
	// A SQL implementation could contain one or more connection strings.
	Nodes []string
	// Database allows multiple isolated stores to be kept in one backend, if supported.
	Database string
	// Collection is analagous to a table in database backends or a key prefix in KV backends
	Collection string
	// Context should contain all implementation specific options, using context.WithValue.
	Context context.Context
}

// Option sets values in Options
type Option func(o *Options)

// Nodes contains the addresses or other connection information of the backing storage.
// For example, an etcd implementation would contain the nodes of the cluster.
// A SQL implementation could contain one or more connection strings.
func Nodes(a ...string) Option {
	return func(o *Options) {
		o.Nodes = a
	}
}

// Database allows multiple isolated stores to be kept in one backend, if supported.
func Database(db string) Option {
	return func(o *Options) {
		o.Database = db
	}
}

// Table is analagous to a table in database backends or a key prefix in KV backends
func Collection(t string) Option {
	return func(o *Options) {
		o.Collection = t
	}
}

// WithContext sets the stores context, for any extra configuration
func WithContext(c context.Context) Option {
	return func(o *Options) {
		o.Context = c
	}
}

// ReadFrom the database and table
func ReadFrom(database, collection string) QueryOption {
	return func(r *QueryOptions) {
		r.Database = database
		r.Collection = collection
	}
}

// ReadLimit limits the number of responses to l
func ReadLimit(l int64) QueryOption {
	return func(r *QueryOptions) {
		r.Limit = l
	}
}

// ReadOffset starts returning responses from o. Use in conjunction with Limit for pagination
func ReadOffset(o int64) QueryOption {
	return func(r *QueryOptions) {
		r.Offset = o
	}
}

// WriteTo the database and table
func WriteTo(database, table string) QueryOption {
	return func(w *QueryOptions) {
		w.Database = database
		w.Collection = table
	}
}

// DeleteFrom the database and table
func DeleteFrom(database, collection string) QueryOption {
	return func(d *QueryOptions) {
		d.Database = database
		d.Collection = collection
	}
}
