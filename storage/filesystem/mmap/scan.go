//go:build darwin || linux

package mmap

import (
	"errors"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
)

var (
	ErrObjectNotFound   = errors.New("object not found")
	ErrOffsetNotFound   = errors.New("offset not found in packfile")
	ErrHashParseFailed  = errors.New("failed to parse hash")
	ErrCorruptedIdx     = errors.New("corrupted idx file")
	ErrNilFile          = errors.New("cannot open mmap: file is nil")
	ErrNoFileDescriptor = errors.New("fs does not support access to file descriptor")
)

type PackScanner struct {
	hashSize int
	count    int

	fanoutStart  int
	namesStart   int
	crcStart     int
	off32Start   int
	off64Start   int
	trailerStart int

	packMmap    []byte
	packCleanup func() error
	idxMmap     []byte
	idxCleanup  func() error
	revMmap     []byte
	revCleanup  func() error
}

func NewPackScanner(hashSize int, pack, idx, rev billy.File) (*PackScanner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PackScanner) FindOffset(h plumbing.ObjectID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *PackScanner) FindHash(offset uint64) (plumbing.ObjectID, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectID), nil
}

func (s *PackScanner) Get(h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *PackScanner) GetByOffset(offset uint64) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *PackScanner) getObject(h plumbing.Hash, offset uint64) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *PackScanner) lookupOffset(want uint64) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (s *PackScanner) offset(pos int) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *PackScanner) fanoutEntry(i int) uint32 { _ = "STUB: not implemented"; return 0 }

func searchObjectID(names []byte, left, right int, want plumbing.ObjectID) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func compareObjectID(names []byte, idx int, want []byte) int { _ = "STUB: not implemented"; return 0 }

func (s *PackScanner) Close() error { _ = "STUB: not implemented"; return nil }
