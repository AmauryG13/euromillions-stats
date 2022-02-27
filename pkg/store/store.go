package store

import (
	"errors"
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
	// Create write a data interface and returns a Result
	Create(value interface{}, opts ...CreateOption) (*Result, error)
	// Read takes a filter and optional ReadOptions. It returns matching *Record or an error.
	Read(filter interface{}, opts ...ReadOption) (*Result, error)
	// Update write the changes to the filtered data, and returns an error if the record was not written.
	Update(filter interface{}, changes interface{}, opts ...UpdateOption) (*Result, error)
	// Delete removes the filtered record
	Delete(filter interface{}, opts ...DeleteOption) (*Result, error)
	// Close the store
	Close() error
	// String returns the name of the implementation.
	String() string
}

// Results
type Result struct {
	// Results of Create, Read, Update, Delete operation
	IDs          []interface{}
	AffectedRows int

	// Field added for the Read operation
	Data []interface{}
}
