package transactional

import (
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type IndexStorage struct {
	storer.IndexStorer
	temporal storer.IndexStorer

	set bool
}

func NewIndexStorage(s, temporal storer.IndexStorer) *IndexStorage {
	_ = "STUB: not implemented"
	return nil
}

func (s *IndexStorage) SetIndex(idx *index.Index) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *IndexStorage) Index() (*index.Index, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *IndexStorage) Commit() error { _ = "STUB: not implemented"; return nil }
