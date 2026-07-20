package dotgit

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

var _ (plumbing.EncodedObject) = &EncodedObject{}

type EncodedObject struct {
	dir *DotGit
	h   plumbing.Hash
	t   plumbing.ObjectType
	sz  int64
}

func (e *EncodedObject) Hash() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (e *EncodedObject) Reader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (e *EncodedObject) SetType(plumbing.ObjectType) { _ = "STUB: not implemented"; return }

func (e *EncodedObject) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (e *EncodedObject) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EncodedObject) SetSize(int64) { _ = "STUB: not implemented"; return }

func (e *EncodedObject) Writer() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func NewEncodedObject(dir *DotGit, h plumbing.Hash, t plumbing.ObjectType, size int64) *EncodedObject {
	_ = "STUB: not implemented"
	return nil
}
