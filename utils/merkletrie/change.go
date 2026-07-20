package merkletrie

import (
	"errors"

	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

var ErrEmptyFileName = errors.New("empty filename in tree entry")

type Action int

const (
	_ Action = iota
	Insert
	Delete
	Modify
)

func (a Action) String() string { _ = "STUB: not implemented"; return "" }

type Change struct {
	From noder.Path

	To noder.Path
}

func (c *Change) Action() (Action, error) { _ = "STUB: not implemented"; return *new(Action), nil }

func NewInsert(n noder.Path) Change { _ = "STUB: not implemented"; return *new(Change) }

func NewDelete(n noder.Path) Change { _ = "STUB: not implemented"; return *new(Change) }

func NewModify(a, b noder.Path) Change { _ = "STUB: not implemented"; return *new(Change) }

func (c Change) String() string { _ = "STUB: not implemented"; return "" }

type Changes []Change

func NewChanges() Changes { _ = "STUB: not implemented"; return *new(Changes) }

func (l *Changes) Add(c Change) { _ = "STUB: not implemented"; return }

func (l *Changes) AddRecursiveInsert(root noder.Path) error { _ = "STUB: not implemented"; return nil }

func (l *Changes) AddRecursiveDelete(root noder.Path) error { _ = "STUB: not implemented"; return nil }

type noderToChangeFn func(noder.Path) Change

func (l *Changes) addRecursive(root noder.Path, ctor noderToChangeFn) error {
	_ = "STUB: not implemented"
	return nil
}
