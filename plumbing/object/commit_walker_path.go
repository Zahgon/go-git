package object

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type commitPathIter struct {
	pathFilter    func(string) bool
	sourceIter    CommitIter
	currentCommit *Commit
	checkParent   bool
}

func NewCommitPathIterFromIter(pathFilter func(string) bool, commitIter CommitIter, checkParent bool) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func NewCommitFileIterFromIter(fileName string, commitIter CommitIter, checkParent bool) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (c *commitPathIter) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *commitPathIter) getNextFileCommit() (*Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *commitPathIter) hasFileChange(changes Changes, parent *Commit) bool {
	_ = "STUB: not implemented"
	return false
}

func isParentHash(hash plumbing.Hash, commit *Commit) bool { _ = "STUB: not implemented"; return false }

func (c *commitPathIter) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *commitPathIter) Close() { _ = "STUB: not implemented"; return }
