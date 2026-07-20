package object

import (
	"context"

	"github.com/go-git/go-git/v6/utils/merkletrie"
)

type Change struct {
	From ChangeEntry
	To   ChangeEntry
}

var empty ChangeEntry

func (c *Change) Action() (merkletrie.Action, error) {
	_ = "STUB: not implemented"
	return *new(merkletrie.Action), nil
}

func (c *Change) Files() (from, to *File, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Change) String() string { _ = "STUB: not implemented"; return "" }

func (c *Change) Patch() (*Patch, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Change) PatchContext(ctx context.Context) (*Patch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Change) name() string { _ = "STUB: not implemented"; return "" }

type ChangeEntry struct {
	Name string

	Tree *Tree

	TreeEntry TreeEntry
}

type Changes []*Change

func (c Changes) Len() int { _ = "STUB: not implemented"; return 0 }

func (c Changes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (c Changes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (c Changes) String() string { _ = "STUB: not implemented"; return "" }

func (c Changes) Patch() (*Patch, error) { _ = "STUB: not implemented"; return nil, nil }

func (c Changes) PatchContext(ctx context.Context) (*Patch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
