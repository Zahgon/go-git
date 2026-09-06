package object

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type Blob struct {
	Hash plumbing.Hash

	Size int64

	obj plumbing.EncodedObject
}

func GetBlob(s storer.EncodedObjectStorer, h plumbing.Hash) (*Blob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodeBlob(o plumbing.EncodedObject) (*Blob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Blob) ID() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (b *Blob) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (b *Blob) Decode(o plumbing.EncodedObject) error { _ = "STUB: not implemented"; return nil }

func (b *Blob) Encode(o plumbing.EncodedObject) (err error) { _ = "STUB: not implemented"; return nil }

func (b *Blob) Reader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type BlobIter struct {
	storer.EncodedObjectIter
	s storer.EncodedObjectStorer
}

func NewBlobIter(s storer.EncodedObjectStorer, iter storer.EncodedObjectIter) *BlobIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *BlobIter) Next() (*Blob, error) { _ = "STUB: not implemented"; return nil, nil }

func (iter *BlobIter) ForEach(cb func(*Blob) error) error { _ = "STUB: not implemented"; return nil }
