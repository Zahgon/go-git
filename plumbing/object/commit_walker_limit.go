package object

import (
	"time"

	"github.com/go-git/go-git/v6/plumbing"
)

type commitLimitIter struct {
	sourceIter   CommitIter
	limitOptions LogLimitOptions
}

type LogLimitOptions struct {
	Since    *time.Time
	Until    *time.Time
	TailHash plumbing.Hash
}

func NewCommitLimitIterFromIter(commitIter CommitIter, limitOptions LogLimitOptions) CommitIter {
	_ = "STUB: not implemented"
	return *new(CommitIter)
}

func (c *commitLimitIter) Next() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *commitLimitIter) ForEach(cb func(*Commit) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *commitLimitIter) Close() { _ = "STUB: not implemented"; return }
