package store

import (
	"errors"
	"time"
)

var (
	// ErrNotFound is returned when a key doesn't exist
	ErrNotFound = errors.New("not found")
)

// Store is a data storage interface
type Store interface {
	// Init initialises the store. It must perform any required setup on the backing storage implementation and check that it is ready for use, returning any errors.
	Init(...Option) error
	// Options allows you to view the current options.
	Options() Options
	// Create writes one (or multiple) query(s) to the store and return an error if creating fails.
	Create(q []*Query, opts ...QueryOption) ([]*Record, error)
	// Read find one query from the store. It returns an array of records and a possible error
	Read(q []*Query, opts ...QueryOption) ([]*Record, error)
	// Update write changes of one (or multiple) query(ies) and return an error if updating fails.
	Update(q []*Query, opts ...WriteOption) ([]*Record, error)
	// Delete removes the record with the corresponding key from the store.
	Delete(q []*Query, opts ...DeleteOption) ([]*Record, error)
	// Close the store
	Close() error
	// String returns the name of the implementation.
	String() string
}

// QueryOptions is an optional option to the query
type QueryOptions struct {
	Database, Collection string
	// Fields lists the fields that needs to be included in the record
	Fields []byte
	// Limit limits the number of returned records
	Limit uint
	// Offset when combined with Limit supports pagination
	Offset uint
	// Sort sorts the retured records
	Sort []byte
}

type QueryOption func(o *QueryOptions)

// Query is an item passed to a Store in order to be stored
type Query struct {
	// The filter to query
	Filter []byte `json:filter`
	// The document to query
	Doc []byte `json:doc`
	// The additional options to the query
	Options QueryOptions
}

// Record is an item stored or retrieved from a Store
type Record struct {
	// The key to store the record
	Key string `json:"key"`
	// The value within the record
	Value []byte `json:"value"`
	// Any associated metadata for indexing
	Metadata map[string]interface{} `json:"metadata"`
	// Time to expire a record: TODO: change to timestamp
	Expiry time.Duration `json:"expiry,omitempty"`
}
