package soiree

import (
	"fmt"

	"github.com/datumforge/datum/pkg/utils/ulids"
)

// EmitterOption defines a function type for Soiree configuration options
type EmitterOption func(Soiree)

var DefaultErrorHandler = func(event Event, err error) error {
	return err
}

// DefaultIDGenerator generates a unique identifier
var DefaultIDGenerator = func() string {
	return ulids.New().String()
}

// DefaultPanicHandler handles panics by printing the panic value
var DefaultPanicHandler = func(p interface{}) {
	fmt.Printf("Panic occurred: %v\n", p)
}

// WithErrorHandler sets a custom error handler for an Soiree
func WithErrorHandler(errHandler func(Event, error) error) EmitterOption {
	return func(m Soiree) {
		m.SetErrorHandler(errHandler)
	}
}

// WithIDGenerator sets a custom ID generator for an Soiree
func WithIDGenerator(idGen func() string) EmitterOption {
	return func(m Soiree) {
		m.SetIDGenerator(idGen)
	}
}

// WithPool sets a custom pool for an Soiree
func WithPool(pool Pool) EmitterOption {
	return func(m Soiree) {
		m.SetPool(pool)
	}
}

// PanicHandler is a function type that handles panics
type PanicHandler func(interface{})

// WithPanicHandler sets a custom panic handler for an Soiree
func WithPanicHandler(panicHandler PanicHandler) EmitterOption {
	return func(m Soiree) {
		m.SetPanicHandler(panicHandler)
	}
}

// WithErrChanBufferSize sets the size of the buffered channel for errors returned by asynchronous emits
func WithErrChanBufferSize(size int) EmitterOption {
	return func(m Soiree) {
		m.SetErrChanBufferSize(size)
	}
}
