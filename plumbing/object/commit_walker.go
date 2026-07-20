package object

import (
	"container/list"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/storage"
)

type commitPreIterator struct {
	seenExternal map[plumbing.Hash]bool
	seen         map[plumbing.Hash]bool
	stack        []CommitIter
	start        *Commit
}

func forEachCommit(next func() (*Commit, error), cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCommitPreorderIter(
	c *Commit,
	seenExternal map[plumbing.Hash]bool,
	ignore []plumbing.Hash,
) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (w *commitPreIterator) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func filteredParentIter(c *Commit, seen map[plumbing.Hash]bool) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (w *commitPreIterator) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *commitPreIterator) Close() { _ = "STUB: not implemented"; return }

type commitPostIterator struct {
	stack []*Commit
	seen  map[plumbing.Hash]bool
}

func NewCommitPostorderIter(c *Commit, ignore []plumbing.Hash) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (w *commitPostIterator) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *commitPostIterator) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *commitPostIterator) Close() { _ = "STUB: not implemented"; return }

type commitPostIteratorFirstParent struct {
	stack []*Commit
	seen  map[plumbing.Hash]bool
}

func NewCommitPostorderIterFirstParent(c *Commit, ignore []plumbing.Hash) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (w *commitPostIteratorFirstParent) Next() (*Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *commitPostIteratorFirstParent) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *commitPostIteratorFirstParent) Close() { _ = "STUB: not implemented"; return }

type commitAllIterator struct {
	currCommit *list.Element
}

func NewCommitAllIter(repoStorer storage.Storer, commitIterFunc func(*Commit) CommitIter) (CommitIter, error) {
	_ = "STUB: not implemented"
	return *new(CommitIter), nil
}

func addReference(
	repoStorer storage.Storer,
	commitIterFunc func(*Commit) CommitIter,
	ref *plumbing.Reference,
	commitsPath *list.List,
	commitsLookup map[plumbing.Hash]*list.Element,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *commitAllIterator) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (it *commitAllIterator) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *commitAllIterator) Close() { _ = "STUB: not implemented"; return }
