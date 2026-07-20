package transactional

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/reflog"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type ReflogStorage struct {
	base     storer.ReflogStorer
	temporal storer.ReflogStorer

	appended map[plumbing.ReferenceName]struct{}

	deleted map[plumbing.ReferenceName]struct{}
}

func NewReflogStorage(base, temporal storer.ReflogStorer) *ReflogStorage {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReflogStorage) Reflog(name plumbing.ReferenceName) ([]*reflog.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ReflogStorage) AppendReflog(name plumbing.ReferenceName, entry *reflog.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReflogStorage) DeleteReflog(name plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReflogStorage) Commit() error { _ = "STUB: not implemented"; return nil }
