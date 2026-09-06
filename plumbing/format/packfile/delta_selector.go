package packfile

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

const (
	maxDepth = int64(50)
)

var applyDelta = map[plumbing.ObjectType]bool{
	plumbing.BlobObject: true,
	plumbing.TreeObject: true,
}

type DeltaSelector struct {
	storer storer.EncodedObjectStorer
}

func NewDeltaSelector(s storer.EncodedObjectStorer) *DeltaSelector {
	_ = "STUB: not implemented"
	return nil
}

func (dw *DeltaSelector) ObjectsToPack(
	hashes []plumbing.Hash,
	packWindow uint,
) ([]*ObjectToPack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dw *DeltaSelector) objectsToPack(
	hashes []plumbing.Hash,
	packWindow uint,
) ([]*ObjectToPack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dw *DeltaSelector) encodedDeltaObject(h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (dw *DeltaSelector) encodedObject(h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (dw *DeltaSelector) fixAndBreakChains(objectsToPack []*ObjectToPack) error {
	_ = "STUB: not implemented"
	return nil
}

func (dw *DeltaSelector) fixAndBreakChainsOne(
	objectsToPack map[plumbing.Hash]*ObjectToPack,
	otp *ObjectToPack,
	visiting map[plumbing.Hash]bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (dw *DeltaSelector) restoreOriginal(otp *ObjectToPack) error {
	_ = "STUB: not implemented"
	return nil
}

func (dw *DeltaSelector) undeltify(otp *ObjectToPack) error { _ = "STUB: not implemented"; return nil }

func (dw *DeltaSelector) sort(objectsToPack []*ObjectToPack) { _ = "STUB: not implemented"; return }

func (dw *DeltaSelector) walk(
	objectsToPack []*ObjectToPack,
	packWindow uint,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (dw *DeltaSelector) tryToDeltify(indexMap map[plumbing.Hash]*deltaIndex, base, target *ObjectToPack) error {
	_ = "STUB: not implemented"
	return nil
}

func (dw *DeltaSelector) deltaSizeLimit(targetSize int64, baseDepth int,
	targetDepth int, targetDelta bool,
) int64 {
	_ = "STUB: not implemented"
	return 0
}

type byTypeAndSize []*ObjectToPack

func (a byTypeAndSize) Len() int { _ = "STUB: not implemented"; return 0 }

func (a byTypeAndSize) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (a byTypeAndSize) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
