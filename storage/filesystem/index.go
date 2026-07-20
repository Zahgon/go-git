package filesystem

import (
	"hash"

	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

type IndexStorage struct {
	dir      *dotgit.DotGit
	h        hash.Hash
	cache    IndexCache
	skipHash bool
}

func (s *IndexStorage) SetIndex(idx *index.Index) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *IndexStorage) writeIndex(idx *index.Index) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *IndexStorage) Index() (i *index.Index, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyIndex(idx *index.Index) *index.Index { _ = "STUB: not implemented"; return nil }
