package transactional

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type ObjectStorage struct {
	storer.EncodedObjectStorer
	temporal storer.EncodedObjectStorer
}

func NewObjectStorage(base, temporal storer.EncodedObjectStorer) *ObjectStorage {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectStorage) SetEncodedObject(obj plumbing.EncodedObject) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (o *ObjectStorage) HasEncodedObject(h plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectStorage) EncodedObjectSize(h plumbing.Hash) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ObjectStorage) EncodedObject(t plumbing.ObjectType, h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (o *ObjectStorage) IterEncodedObjects(t plumbing.ObjectType) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func (o *ObjectStorage) Commit() error { _ = "STUB: not implemented"; return nil }

func (o *ObjectStorage) AddAlternate(remote string) error { _ = "STUB: not implemented"; return nil }
