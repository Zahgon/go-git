package filesystem

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

type ShallowStorage struct {
	dir *dotgit.DotGit
}

func (s *ShallowStorage) SetShallow(commits []plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ShallowStorage) Shallow() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
