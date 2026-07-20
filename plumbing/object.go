package plumbing

import (
	"errors"
	"io"
)

var (
	ErrObjectNotFound = errors.New("object not found")

	ErrInvalidType = errors.New("invalid object type")
)

type EncodedObject interface {
	Hash() Hash
	Type() ObjectType
	SetType(ObjectType)
	Size() int64
	SetSize(int64)
	Reader() (io.ReadCloser, error)
	Writer() (io.WriteCloser, error)
}

type DeltaObject interface {
	EncodedObject

	BaseHash() Hash

	ActualHash() Hash

	ActualSize() int64
}

type ObjectType int8

const (
	InvalidObject ObjectType = 0

	CommitObject ObjectType = 1

	TreeObject ObjectType = 2

	BlobObject ObjectType = 3

	TagObject ObjectType = 4

	OFSDeltaObject ObjectType = 6

	REFDeltaObject ObjectType = 7

	AnyObject ObjectType = -127
)

func (t ObjectType) String() string { _ = "STUB: not implemented"; return "" }

func (t ObjectType) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (t ObjectType) Valid() bool { _ = "STUB: not implemented"; return false }

func (t ObjectType) IsDelta() bool { _ = "STUB: not implemented"; return false }

func ParseObjectType(value string) (typ ObjectType, err error) {
	_ = "STUB: not implemented"
	return *new(ObjectType), nil
}
