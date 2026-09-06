package plugin

import (
	"errors"
	"sync"
)

var (
	ErrFrozen = errors.New("plugin registry is frozen")

	ErrNotFound = errors.New("plugin not found")

	ErrNilFactory = errors.New("factory must not be nil")
)

var (
	mu      sync.RWMutex
	entries = map[Name]*entry{}
)

type Name string

func Register[T any](key key[T], factory func() T) error { _ = "STUB: not implemented"; return nil }

func Get[T any](key key[T]) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func Has[T any](key key[T]) bool { _ = "STUB: not implemented"; return false }

type key[T any] struct {
	name Name
}

func newKey[T any](name Name) key[T] { _ = "STUB: not implemented"; return nil }

func newKeyWithValidator[T any](name Name, validate func(T) error) key[T] {
	_ = "STUB: not implemented"
	return nil
}

type entry struct {
	frozen   bool
	factory  any
	validate func(any) error
}

//nolint:unused // accessed via go:linkname from test files in the root package
func resetEntry(name Name) { _ = "STUB: not implemented"; return }
