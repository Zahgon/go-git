package plumbing

import (
	"bytes"
	"io"
)

func NewMemoryObject(oh *ObjectHasher) *MemoryObject { _ = "STUB: not implemented"; return nil }

type MemoryObject struct {
	t    ObjectType
	h    Hash
	cont []byte
	sz   int64
	oh   *ObjectHasher
}

func (o *MemoryObject) Hash() Hash { _ = "STUB: not implemented"; return *new(Hash) }

func (o *MemoryObject) Type() ObjectType { _ = "STUB: not implemented"; return *new(ObjectType) }

func (o *MemoryObject) SetType(t ObjectType) { _ = "STUB: not implemented"; return }

func (o *MemoryObject) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (o *MemoryObject) SetSize(s int64) { _ = "STUB: not implemented"; return }

func (o *MemoryObject) Reader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (o *MemoryObject) Writer() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (o *MemoryObject) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (o *MemoryObject) Close() error { _ = "STUB: not implemented"; return nil }

type nopCloser struct {
	*bytes.Reader
}

func (nc nopCloser) Close() error { _ = "STUB: not implemented"; return nil }
