package filesystem

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

type ReferenceStorage struct {
	dir *dotgit.DotGit
}

func (r *ReferenceStorage) SetReference(ref *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReferenceStorage) CheckAndSetReference(ref, old *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReferenceStorage) Reference(n plumbing.ReferenceName) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReferenceStorage) IterReferences() (storer.ReferenceIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceIter), nil
}

func (r *ReferenceStorage) RemoveReference(n plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReferenceStorage) CountLooseRefs() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *ReferenceStorage) PackRefs() error { _ = "STUB: not implemented"; return nil }
