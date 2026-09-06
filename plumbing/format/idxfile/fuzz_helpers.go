package idxfile

import (
	"bytes"

	"github.com/go-git/go-git/v6/plumbing"
)

func buildMinimalIdx(count, hashSize int) []byte { _ = "STUB: not implemented"; return nil }

func buildMinimalRev(count, hashSize int) []byte { _ = "STUB: not implemented"; return nil }

func buildOOBOffset64Idx() ([]byte, plumbing.Hash) {
	_ = "STUB: not implemented"
	return nil, *new(plumbing.Hash)
}

type nopCloserReaderAt struct {
	*bytes.Reader
}

func (nopCloserReaderAt) Close() error { _ = "STUB: not implemented"; return nil }
