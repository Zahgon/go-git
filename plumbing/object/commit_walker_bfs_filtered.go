package object

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

func NewFilterCommitIter(
	from *Commit,
	isValid *CommitFilter,
	isLimit *CommitFilter,
) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

type CommitFilter func(*Commit) bool

type filterCommitIter struct {
	isValid CommitFilter
	isLimit CommitFilter
	visited map[plumbing.Hash]struct{}
	queue   []*Commit
	lastErr error
}

func (w *filterCommitIter) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *filterCommitIter) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *filterCommitIter) Error() error { _ = "STUB: not implemented"; return nil }

func (w *filterCommitIter) Close() { _ = "STUB: not implemented"; return }

func (w *filterCommitIter) close(err error) error { _ = "STUB: not implemented"; return nil }

func (w *filterCommitIter) popNewFromQueue() (*Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *filterCommitIter) addToQueue(
	store storer.EncodedObjectStorer,
	hashes ...plumbing.Hash,
) error {
	_ = "STUB: not implemented"
	return nil
}
