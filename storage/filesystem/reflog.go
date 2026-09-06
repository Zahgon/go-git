package filesystem

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/reflog"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

type ReflogStorage struct {
	dir *dotgit.DotGit
}

func (r *ReflogStorage) Reflog(name plumbing.ReferenceName) ([]*reflog.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReflogStorage) AppendReflog(name plumbing.ReferenceName, entry *reflog.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReflogStorage) DeleteReflog(name plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}
