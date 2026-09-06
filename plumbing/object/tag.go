package object

import (
	"errors"

	"github.com/ProtonMail/go-crypto/openpgp"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

var ErrMalformedTag = errors.New("malformed tag")

type Tag struct {
	Hash plumbing.Hash

	Name string

	Tagger Signature

	Message string

	Signature string

	SignatureSHA256 string

	TargetType plumbing.ObjectType

	Target plumbing.Hash

	s storer.EncodedObjectStorer

	src plumbing.EncodedObject
}

func GetTag(s storer.EncodedObjectStorer, h plumbing.Hash) (*Tag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodeTag(s storer.EncodedObjectStorer, o plumbing.EncodedObject) (*Tag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tag) ID() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (t *Tag) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (t *Tag) reset() { _ = "STUB: not implemented"; return }

func (t *Tag) Decode(o plumbing.EncodedObject) (err error) { _ = "STUB: not implemented"; return nil }

func (t *Tag) Encode(o plumbing.EncodedObject) error { _ = "STUB: not implemented"; return nil }

func (t *Tag) EncodeWithoutSignature(o plumbing.EncodedObject) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tag) matchesSource() bool { _ = "STUB: not implemented"; return false }

func (t *Tag) encode(o plumbing.EncodedObject, includeSig bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func isZeroSignature(s Signature) bool { _ = "STUB: not implemented"; return false }

func (t *Tag) Commit() (*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tag) Tree() (*Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tag) Blob() (*Blob, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tag) Object() (Object, error) { _ = "STUB: not implemented"; return *new(Object), nil }

func (t *Tag) String() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) Verify(armoredKeyRing string) (*openpgp.Entity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TagIter struct {
	storer.EncodedObjectIter
	s storer.EncodedObjectStorer
}

func NewTagIter(s storer.EncodedObjectStorer, iter storer.EncodedObjectIter) *TagIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *TagIter) Next() (*Tag, error) { _ = "STUB: not implemented"; return nil, nil }

func (iter *TagIter) ForEach(cb func(*Tag) error) error { _ = "STUB: not implemented"; return nil }

func objectAsString(obj Object) string { _ = "STUB: not implemented"; return "" }
