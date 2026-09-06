package object

import (
	"github.com/emirpasic/gods/trees/binaryheap"

	"github.com/go-git/go-git/v6/plumbing"
)

type commitIteratorByCTime struct {
	seenExternal map[plumbing.Hash]bool
	seen         map[plumbing.Hash]bool
	heap         *binaryheap.Heap
}

func NewCommitIterCTime(
	c *Commit,
	seenExternal map[plumbing.Hash]bool,
	ignore []plumbing.Hash,
) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (w *commitIteratorByCTime) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *commitIteratorByCTime) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *commitIteratorByCTime) Close() { _ = "STUB: not implemented"; return }
