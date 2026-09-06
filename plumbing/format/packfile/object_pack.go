package packfile

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type ObjectToPack struct {
	Object plumbing.EncodedObject

	Base *ObjectToPack

	Original plumbing.EncodedObject

	Depth int

	Offset int64

	resolvedOriginal bool
	originalType     plumbing.ObjectType
	originalSize     int64
	originalHash     plumbing.Hash
}

func newObjectToPack(o plumbing.EncodedObject) *ObjectToPack { _ = "STUB: not implemented"; return nil }

func newDeltaObjectToPack(base *ObjectToPack, original, delta plumbing.EncodedObject) *ObjectToPack {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectToPack) BackToOriginal() { _ = "STUB: not implemented"; return }

func (o *ObjectToPack) IsWritten() bool { _ = "STUB: not implemented"; return false }

func (o *ObjectToPack) MarkWantWrite() { _ = "STUB: not implemented"; return }

func (o *ObjectToPack) WantWrite() bool { _ = "STUB: not implemented"; return false }

func (o *ObjectToPack) SetOriginal(obj plumbing.EncodedObject) { _ = "STUB: not implemented"; return }

func (o *ObjectToPack) SaveOriginalMetadata() { _ = "STUB: not implemented"; return }

func (o *ObjectToPack) CleanOriginal() { _ = "STUB: not implemented"; return }

func (o *ObjectToPack) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (o *ObjectToPack) Hash() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (o *ObjectToPack) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (o *ObjectToPack) IsDelta() bool { _ = "STUB: not implemented"; return false }

func (o *ObjectToPack) SetDelta(base *ObjectToPack, delta plumbing.EncodedObject) {
	_ = "STUB: not implemented"
	return
}
