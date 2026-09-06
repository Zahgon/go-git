package filesystem

import (
	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type lazyPackfilesIter struct {
	hashes []plumbing.Hash
	open   func(h plumbing.Hash) (storer.EncodedObjectIter, error)
	cur    storer.EncodedObjectIter
}

func (it *lazyPackfilesIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (it *lazyPackfilesIter) ForEach(cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *lazyPackfilesIter) Close() { _ = "STUB: not implemented"; return }

type packfileIter struct {
	pack billy.File
	iter storer.EncodedObjectIter
	seen map[plumbing.Hash]struct{}

	keepPack bool
}

func NewPackfileIter(
	fs billy.Filesystem,
	f billy.File,
	idxFile billy.File,
	t plumbing.ObjectType,
	keepPack bool,
	_ int64,
	objectIDSize int,
) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func newPackfileIter(
	fs billy.Filesystem,
	f billy.File,
	t plumbing.ObjectType,
	seen map[plumbing.Hash]struct{},
	index idxfile.Index,
	cache cache.Object,
	keepPack bool,
	objectIDSize int,
) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func (iter *packfileIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (iter *packfileIter) ForEach(cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *packfileIter) Close() { _ = "STUB: not implemented"; return }

type objectsIter struct {
	s *ObjectStorage
	t plumbing.ObjectType
	h []plumbing.Hash
}

func (iter *objectsIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (iter *objectsIter) ForEach(cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *objectsIter) Close() { _ = "STUB: not implemented"; return }
