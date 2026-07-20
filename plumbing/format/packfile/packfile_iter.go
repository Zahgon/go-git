package packfile

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
)

type objectIter struct {
	p    *Packfile
	typ  plumbing.ObjectType
	iter idxfile.EntryIter
}

func (i *objectIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (i *objectIter) next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (i *objectIter) ForEach(f func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *objectIter) Close() { _ = "STUB: not implemented"; return }
