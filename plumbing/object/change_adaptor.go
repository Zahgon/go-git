package object

import (
	"github.com/go-git/go-git/v6/utils/merkletrie"
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

func newChange(c merkletrie.Change) (*Change, error) { _ = "STUB: not implemented"; return nil, nil }

func newChangeEntry(p noder.Path) (ChangeEntry, error) {
	_ = "STUB: not implemented"
	return *new(ChangeEntry), nil
}

func newChanges(src merkletrie.Changes) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}
