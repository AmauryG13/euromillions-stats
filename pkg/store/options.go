package store

import (
	"context"

	"go-micro.dev/v4/client"
)

// Options contains configuration for the Store
type Options struct {
	// Nodes contains the addresses or other connection information of the backing storage.
	// For example, an etcd implementation would contain the nodes of the cluster.
	// A SQL implementation could contain one or more connection strings.
	Nodes []string
	// Database allows multiple isolated stores to be kept in one backend, if supported.
	Database string
	// Collection is analagous to a Collection in database backends or a key prefix in KV backends
	Collection string
	// Context should contain all implementation specific options, using context.WithValue.
	Context context.Context
	// Client to use for RPC
	Client client.Client
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

// Collection is analagous to a Collection in database backends or a key prefix in KV backends
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

// CreateOption configures an individual Create operation
type CreateOptions struct {
	Database, Collection string
}

type CreateOption func(*CreateOptions)

func CreateOn(database, Collection string) CreateOption {
	return func(c *CreateOptions) {
		c.Database = database
		c.Collection = Collection
	}
}

// ReadOptions configures an individual Read operation
type ReadOptions struct {
	Database, Collection string
	// Fields limits the fields that need to be outputed
	Fields []byte
	// Limit limits the number of returned records
	Limit int64
	// Offset when combined with Limit supports pagination
	Offset int64
	// Sort helps sorting the data
	Sort []byte
	// Schema is an interface how to render the read data
	Schema interface{}
}

// ReadOption sets values in ReadOptions
type ReadOption func(r *ReadOptions)

// ReadFrom the database and Collection
func ReadFrom(database, Collection string) ReadOption {
	return func(r *ReadOptions) {
		r.Database = database
		r.Collection = Collection
	}
}

func ReadFields(f []byte) ReadOption {
	return func(r *ReadOptions) {
		r.Fields = f
	}
}

// ReadLimit limits the number of responses to l
func ReadLimit(l int64) ReadOption {
	return func(r *ReadOptions) {
		r.Limit = l
	}
}

// ReadOffset starts returning responses from o. Use in conjunction with Limit for pagination
func ReadOffset(o int64) ReadOption {
	return func(r *ReadOptions) {
		r.Offset = o
	}
}

func ReadFirst() ReadOption {
	return func(r *ReadOptions) {
		r.Limit = 1
		r.Offset = 0
	}
}

func ReadSort(s []byte) ReadOption {
	return func(r *ReadOptions) {
		r.Sort = s
	}
}

func ReadSchema(s interface{}) ReadOption {
	return func(r *ReadOptions) {
		r.Schema = s
	}
}

// UpdateOptions configures an individual Update operation
type UpdateOptions struct {
	Database, Collection string
}

// UpdateOption sets values in UpdateOptions
type UpdateOption func(d *UpdateOptions)

// UpdateFrom the database and Collection
func UpdateOn(database, Collection string) UpdateOption {
	return func(d *UpdateOptions) {
		d.Database = database
		d.Collection = Collection
	}
}

// DeleteOptions configures an individual Delete operation
type DeleteOptions struct {
	Database, Collection string
}

// DeleteOption sets values in DeleteOptions
type DeleteOption func(d *DeleteOptions)

// DeleteFrom the database and Collection
func DeleteFrom(database, Collection string) DeleteOption {
	return func(d *DeleteOptions) {
		d.Database = database
		d.Collection = Collection
	}
}
