package nosql

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
	// Create writes one (or multiple) record(s) to the store and return an error if creating fails.
	Create(r []*Record, opts ...WriteOption) error
	// Read extracts one (or multiple) record(s) from the store. It returns an array of records and a possible error
	Read(r []*Record, opts ...ReadOption) ([]*Record, error)
	// Update write changes to one (or multiple) records and return an error if updating fails.
	Update(r []*Record, opts ...WriteOption) error
	// Delete removes the record with the corresponding key from the store.
	Delete(r []*Record, opts ...DeleteOption) error
	// Close the store
	Close() error
	// String returns the name of the implementation.
	String() string
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
