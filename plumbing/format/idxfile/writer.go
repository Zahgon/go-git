package idxfile

import (
	"sync"

	"github.com/go-git/go-git/v6/plumbing"
)

type objects []Entry

type Writer struct {
	m sync.Mutex

	count    uint32
	checksum plumbing.Hash
	objects  objects
	offset64 uint32
	finished bool
	index    *MemoryIndex
	added    map[plumbing.Hash]struct{}
}

func (w *Writer) Index() (*MemoryIndex, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Writer) Add(h plumbing.Hash, pos uint64, crc uint32) { _ = "STUB: not implemented"; return }

func (w *Writer) Finished() bool { _ = "STUB: not implemented"; return false }

func (w *Writer) OnHeader(count uint32) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) OnInflatedObjectHeader(_ plumbing.ObjectType, _, _ int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) OnInflatedObjectContent(h plumbing.Hash, pos int64, crc uint32, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) OnFooter(h plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) createIndex() (*MemoryIndex, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Writer) addOffset64(pos uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o objects) Len() int { _ = "STUB: not implemented"; return 0 }

func (o objects) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (o objects) Swap(i, j int) { _ = "STUB: not implemented"; return }
