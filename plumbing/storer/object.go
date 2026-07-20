package storer

import (
	"errors"
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
)

var ErrStop = errors.New("stop iter")

type EncodedObjectStorer interface {
	RawObjectWriter(typ plumbing.ObjectType, sz int64) (w io.WriteCloser, err error)

	NewEncodedObject() plumbing.EncodedObject

	SetEncodedObject(plumbing.EncodedObject) (plumbing.Hash, error)

	EncodedObject(plumbing.ObjectType, plumbing.Hash) (plumbing.EncodedObject, error)

	IterEncodedObjects(plumbing.ObjectType) (EncodedObjectIter, error)

	HasEncodedObject(plumbing.Hash) error

	EncodedObjectSize(plumbing.Hash) (int64, error)
	AddAlternate(remote string) error
}

type DeltaObjectStorer interface {
	DeltaObject(plumbing.ObjectType, plumbing.Hash) (plumbing.EncodedObject, error)
}

type Transactioner interface {
	Begin() Transaction
}

type LooseObjectStorer interface {
	ForEachObjectHash(func(plumbing.Hash) error) error

	LooseObjectTime(plumbing.Hash) (time.Time, error)

	DeleteLooseObject(plumbing.Hash) error
}

type PackedObjectStorer interface {
	ObjectPacks() ([]plumbing.Hash, error)

	DeleteOldObjectPackAndIndex(plumbing.Hash, time.Time) error
}

type PackfileWriter interface {
	PackfileWriter() (io.WriteCloser, error)
}

type EncodedObjectIter interface {
	Next() (plumbing.EncodedObject, error)
	ForEach(func(plumbing.EncodedObject) error) error
	Close()
}

type Transaction interface {
	SetEncodedObject(plumbing.EncodedObject) (plumbing.Hash, error)
	EncodedObject(plumbing.ObjectType, plumbing.Hash) (plumbing.EncodedObject, error)
	Commit() error
	Rollback() error
}

type EncodedObjectLookupIter struct {
	storage EncodedObjectStorer
	series  []plumbing.Hash
	t       plumbing.ObjectType
	pos     int
}

func NewEncodedObjectLookupIter(
	storage EncodedObjectStorer, t plumbing.ObjectType, series []plumbing.Hash,
) *EncodedObjectLookupIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *EncodedObjectLookupIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (iter *EncodedObjectLookupIter) ForEach(cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *EncodedObjectLookupIter) Close() { _ = "STUB: not implemented"; return }

type EncodedObjectSliceIter struct {
	series []plumbing.EncodedObject
}

func NewEncodedObjectSliceIter(series []plumbing.EncodedObject) *EncodedObjectSliceIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *EncodedObjectSliceIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (iter *EncodedObjectSliceIter) ForEach(cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *EncodedObjectSliceIter) Close() { _ = "STUB: not implemented"; return }

type MultiEncodedObjectIter struct {
	iters []EncodedObjectIter
}

func NewMultiEncodedObjectIter(iters []EncodedObjectIter) EncodedObjectIter {
	_ = "STUB: not implemented"
	return *new(EncodedObjectIter)
}

func (iter *MultiEncodedObjectIter) Next() (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (iter *MultiEncodedObjectIter) ForEach(cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *MultiEncodedObjectIter) Close() { _ = "STUB: not implemented"; return }

type bareIterator interface {
	Next() (plumbing.EncodedObject, error)
	Close()
}

func ForEachIterator(iter bareIterator, cb func(plumbing.EncodedObject) error) error {
	_ = "STUB: not implemented"
	return nil
}
