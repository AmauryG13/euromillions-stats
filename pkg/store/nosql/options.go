package nosql

import (
	"context"
	"time"
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

// ReadOptions configures an individual Read operation
type ReadOptions struct {
	Database, Collection string
	// Limit limits the number of returned records
	Limit uint
	// Offset when combined with Limit supports pagination
	Offset uint
}

// ReadOption sets values in ReadOptions
type ReadOption func(r *ReadOptions)

// ReadFrom the database and table
func ReadFrom(database, collection string) ReadOption {
	return func(r *ReadOptions) {
		r.Database = database
		r.Collection = collection
	}
}

// ReadLimit limits the number of responses to l
func ReadLimit(l uint) ReadOption {
	return func(r *ReadOptions) {
		r.Limit = l
	}
}

// ReadOffset starts returning responses from o. Use in conjunction with Limit for pagination
func ReadOffset(o uint) ReadOption {
	return func(r *ReadOptions) {
		r.Offset = o
	}
}

// WriteOptions configures an individual Write operation
// If Expiry and TTL are set TTL takes precedence
type WriteOptions struct {
	Database, Collection string
	// Expiry is the time the record expires
	Expiry time.Time
	// TTL is the time until the record expires
	TTL time.Duration
}

// WriteOption sets values in WriteOptions
type WriteOption func(w *WriteOptions)

// WriteTo the database and table
func WriteTo(database, table string) WriteOption {
	return func(w *WriteOptions) {
		w.Database = database
		w.Collection = table
	}
}

// WriteExpiry is the time the record expires
func WriteExpiry(t time.Time) WriteOption {
	return func(w *WriteOptions) {
		w.Expiry = t
	}
}

// WriteTTL is the time the record expires
func WriteTTL(d time.Duration) WriteOption {
	return func(w *WriteOptions) {
		w.TTL = d
	}
}

// DeleteOptions configures an individual Delete operation
type DeleteOptions struct {
	Database, Collection string
}

// DeleteOption sets values in DeleteOptions
type DeleteOption func(d *DeleteOptions)

// DeleteFrom the database and table
func DeleteFrom(database, collection string) DeleteOption {
	return func(d *DeleteOptions) {
		d.Database = database
		d.Collection = collection
	}
}
