package filesystem

import (
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/storage"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

type ModuleStorage struct {
	dir          *dotgit.DotGit
	objectFormat formatcfg.ObjectFormat
}

func (s *ModuleStorage) Module(name string) (storage.Storer, error) {
	_ = "STUB: not implemented"
	return *new(storage.Storer), nil
}
