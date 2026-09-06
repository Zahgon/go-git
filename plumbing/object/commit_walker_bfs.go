package object

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type bfsCommitIterator struct {
	seenExternal map[plumbing.Hash]bool
	seen         map[plumbing.Hash]bool
	queue        []*Commit
}

func NewCommitIterBSF(
	c *Commit,
	seenExternal map[plumbing.Hash]bool,
	ignore []plumbing.Hash,
) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (w *bfsCommitIterator) appendHash(store storer.EncodedObjectStorer, h plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *bfsCommitIterator) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *bfsCommitIterator) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *bfsCommitIterator) Close() { _ = "STUB: not implemented"; return }
