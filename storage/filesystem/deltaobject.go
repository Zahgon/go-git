package filesystem

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type deltaObject struct {
	plumbing.EncodedObject
	base plumbing.Hash
	hash plumbing.Hash
	size int64
}

func newDeltaObject(
	obj plumbing.EncodedObject,
	hash plumbing.Hash,
	base plumbing.Hash,
	size int64,
) plumbing.DeltaObject {
	_ = "STUB: not implemented"
	return *new(plumbing.DeltaObject)
}

func (o *deltaObject) BaseHash() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}

func (o *deltaObject) ActualSize() int64 { _ = "STUB: not implemented"; return 0 }

func (o *deltaObject) ActualHash() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}
