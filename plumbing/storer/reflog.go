package storer

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/reflog"
)

type ReflogStorer interface {
	Reflog(name plumbing.ReferenceName) ([]*reflog.Entry, error)

	AppendReflog(name plumbing.ReferenceName, entry *reflog.Entry) error

	DeleteReflog(name plumbing.ReferenceName) error
}
