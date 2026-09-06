package worktree

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type options struct {
	commit       plumbing.Hash
	detachedHead bool
}

func (o *options) Validate() error { _ = "STUB: not implemented"; return nil }

type Option func(*options)

func WithCommit(commit plumbing.Hash) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDetachedHead() Option { _ = "STUB: not implemented"; return *new(Option) }
