package merkletrie

import (
	"github.com/go-git/go-git/v6/utils/merkletrie/internal/frame"
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

type Iter struct {
	hasStarted bool

	frameStack []*frame.Frame

	base noder.Path
}

func NewIter(n noder.Noder) (*Iter, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIterFromPath(p noder.Path) (*Iter, error) { _ = "STUB: not implemented"; return nil, nil }

func newIter(root noder.Noder, base noder.Path) (*Iter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (iter *Iter) top() (*frame.Frame, bool) { _ = "STUB: not implemented"; return nil, false }

func (iter *Iter) push(f *frame.Frame) { _ = "STUB: not implemented"; return }

const (
	doDescend   = true
	dontDescend = false
)

func (iter *Iter) Next() (noder.Path, error) {
	_ = "STUB: not implemented"
	return *new(noder.Path), nil
}

func (iter *Iter) Step() (noder.Path, error) {
	_ = "STUB: not implemented"
	return *new(noder.Path), nil
}

func (iter *Iter) advance(wantDescend bool) (noder.Path, error) {
	_ = "STUB: not implemented"
	return *new(noder.Path), nil
}

func (iter *Iter) current() (noder.Path, error) {
	_ = "STUB: not implemented"
	return *new(noder.Path), nil
}

func (iter *Iter) drop() { _ = "STUB: not implemented"; return }
