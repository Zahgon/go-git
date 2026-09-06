package storer

import (
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
)

const MaxResolveRecursion = 1024

var ErrMaxResolveRecursion = errors.New("max. recursion level reached")

type ReferenceStorer interface {
	SetReference(*plumbing.Reference) error

	CheckAndSetReference(newRef, old *plumbing.Reference) error
	Reference(plumbing.ReferenceName) (*plumbing.Reference, error)
	IterReferences() (ReferenceIter, error)
	RemoveReference(plumbing.ReferenceName) error
	CountLooseRefs() (int, error)
	PackRefs() error
}

type ReferenceIter interface {
	Next() (*plumbing.Reference, error)
	ForEach(func(*plumbing.Reference) error) error
	Close()
}

type referenceFilteredIter struct {
	ff   func(r *plumbing.Reference) bool
	iter ReferenceIter
}

func NewReferenceFilteredIter(
	ff func(r *plumbing.Reference) bool, iter ReferenceIter,
) ReferenceIter {
	_ = "STUB: not implemented"
	return *new(ReferenceIter)
}

func (iter *referenceFilteredIter) Next() (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (iter *referenceFilteredIter) ForEach(cb func(*plumbing.Reference) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *referenceFilteredIter) Close() { _ = "STUB: not implemented"; return }

type ReferenceSliceIter struct {
	series []*plumbing.Reference
	pos    int
}

func NewReferenceSliceIter(series []*plumbing.Reference) ReferenceIter {
	_ = "STUB: not implemented"
	return *new(ReferenceIter)
}

func (iter *ReferenceSliceIter) Next() (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (iter *ReferenceSliceIter) ForEach(cb func(*plumbing.Reference) error) error {
	_ = "STUB: not implemented"
	return nil
}

type bareReferenceIterator interface {
	Next() (*plumbing.Reference, error)
	Close()
}

func forEachReferenceIter(iter bareReferenceIterator, cb func(*plumbing.Reference) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *ReferenceSliceIter) Close() { _ = "STUB: not implemented"; return }

type MultiReferenceIter struct {
	iters []ReferenceIter
}

func NewMultiReferenceIter(iters []ReferenceIter) ReferenceIter {
	_ = "STUB: not implemented"
	return *new(ReferenceIter)
}

func (iter *MultiReferenceIter) Next() (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (iter *MultiReferenceIter) ForEach(cb func(*plumbing.Reference) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *MultiReferenceIter) Close() { _ = "STUB: not implemented"; return }

func ResolveReference(s ReferenceStorer, n plumbing.ReferenceName) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveReference(s ReferenceStorer, r *plumbing.Reference, recursion int) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
