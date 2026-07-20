package transactional

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type ShallowStorage struct {
	storer.ShallowStorer
	temporal storer.ShallowStorer
}

func NewShallowStorage(base, temporal storer.ShallowStorer) *ShallowStorage {
	_ = "STUB: not implemented"
	return nil
}

func (s *ShallowStorage) SetShallow(commits []plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ShallowStorage) Shallow() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShallowStorage) Commit() error { _ = "STUB: not implemented"; return nil }
