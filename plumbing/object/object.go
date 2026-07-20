package object

import (
	"errors"
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

var ErrUnsupportedObject = errors.New("unsupported object type")

type Object interface {
	ID() plumbing.Hash
	Type() plumbing.ObjectType
	Decode(plumbing.EncodedObject) error
	Encode(plumbing.EncodedObject) error
}

func GetObject(s storer.EncodedObjectStorer, h plumbing.Hash) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}

func DecodeObject(s storer.EncodedObjectStorer, o plumbing.EncodedObject) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}

const DateFormat = "Mon Jan 02 15:04:05 2006 -0700"

type Signature struct {
	Name string

	Email string

	When time.Time
}

func (s *Signature) Decode(b []byte) { _ = "STUB: not implemented"; return }

func (s *Signature) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

var timeZoneLength = 5

func (s *Signature) decodeTimeAndTimeZone(b []byte) { _ = "STUB: not implemented"; return }

func (s *Signature) encodeTimeAndTimeZone(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (s *Signature) String() string { _ = "STUB: not implemented"; return "" }

type ObjectIter struct { //nolint:revive // stutters but is a well-established name
	storer.EncodedObjectIter
	s storer.EncodedObjectStorer
}

func NewObjectIter(s storer.EncodedObjectStorer, iter storer.EncodedObjectIter) *ObjectIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *ObjectIter) Next() (Object, error) { _ = "STUB: not implemented"; return *new(Object), nil }

func (iter *ObjectIter) ForEach(cb func(Object) error) error { _ = "STUB: not implemented"; return nil }

func (iter *ObjectIter) toObject(obj plumbing.EncodedObject) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}
