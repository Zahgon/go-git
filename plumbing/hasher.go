package plumbing

import (
	"hash"
	"sync"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

type ObjectHasher struct {
	hasher hash.Hash
	m      sync.Mutex
	format format.ObjectFormat
}

func (h *ObjectHasher) Size() int { _ = "STUB: not implemented"; return 0 }

func (h *ObjectHasher) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (h *ObjectHasher) Compute(ot ObjectType, d []byte) (ObjectID, error) {
	_ = "STUB: not implemented"
	return *new(ObjectID), nil
}

func FromObjectFormat(f format.ObjectFormat) *ObjectHasher { _ = "STUB: not implemented"; return nil }

func FromHash(h hash.Hash) (*ObjectHasher, error) { _ = "STUB: not implemented"; return nil, nil }

func writeHeader(h hash.Hash, ot ObjectType, sz int64) { _ = "STUB: not implemented"; return }
