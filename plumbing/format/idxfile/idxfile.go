package idxfile

import (
	"sync"

	"github.com/go-git/go-git/v6/plumbing"
)

const (
	VersionSupported = 2

	noMapping = -1
)

var idxHeader = []byte{255, 't', 'O', 'c'}

type Index interface {
	Contains(h plumbing.Hash) (bool, error)

	FindOffset(h plumbing.Hash) (int64, error)

	FindCRC32(h plumbing.Hash) (uint32, error)

	FindHash(o int64) (plumbing.Hash, error)

	Count() (int64, error)

	Entries() (EntryIter, error)

	EntriesByOffset() (EntryIter, error)

	EntriesWithPrefix(prefix []byte) (EntryIter, error)

	MayContain(h plumbing.Hash) bool

	Close() error
}

type MemoryIndex struct {
	Version uint32

	Fanout [256]uint32

	FanoutMapping [256]int

	Names [][]byte

	Offset32 [][]byte

	CRC32 [][]byte

	Offset64 []byte

	PackfileChecksum plumbing.Hash

	IdxChecksum plumbing.Hash

	offsetHash      map[int64]plumbing.Hash
	offsetBuildOnce sync.Once
	mu              sync.RWMutex

	objectIDSize int
}

var _ Index = (*MemoryIndex)(nil)

func (idx *MemoryIndex) Close() error { _ = "STUB: not implemented"; return nil }

func NewMemoryIndex(objectIDSize int) *MemoryIndex { _ = "STUB: not implemented"; return nil }

func (idx *MemoryIndex) findHashIndex(h plumbing.Hash) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (idx *MemoryIndex) MayContain(h plumbing.Hash) bool { _ = "STUB: not implemented"; return false }

func (idx *MemoryIndex) Contains(h plumbing.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (idx *MemoryIndex) FindOffset(h plumbing.Hash) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

const isO64Mask = uint64(1) << 31

func (idx *MemoryIndex) getOffset(firstLevel, secondLevel int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (idx *MemoryIndex) FindCRC32(h plumbing.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (idx *MemoryIndex) getCRC32(firstLevel, secondLevel int) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func (idx *MemoryIndex) FindHash(o int64) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (idx *MemoryIndex) genOffsetHash() error { _ = "STUB: not implemented"; return nil }

func (idx *MemoryIndex) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (idx *MemoryIndex) Entries() (EntryIter, error) {
	_ = "STUB: not implemented"
	return *new(EntryIter), nil
}

func (idx *MemoryIndex) EntriesWithPrefix(prefix []byte) (EntryIter, error) {
	_ = "STUB: not implemented"
	return *new(EntryIter), nil
}

func (idx *MemoryIndex) EntriesByOffset() (EntryIter, error) {
	_ = "STUB: not implemented"
	return *new(EntryIter), nil
}

func (idx *MemoryIndex) idSize() int { _ = "STUB: not implemented"; return 0 }

type EntryIter interface {
	Next() (*Entry, error)

	Close() error
}

type idxfileEntryIter struct {
	idx                     *MemoryIndex
	total                   int
	firstLevel, secondLevel int
}

func (i *idxfileEntryIter) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *idxfileEntryIter) Close() error { _ = "STUB: not implemented"; return nil }

type idxfilePrefixIter struct {
	idSize   int
	prefix   []byte
	names    []byte
	offset32 []byte
	crc32    []byte
	offset64 []byte
	pos      int
	done     bool
}

func (i *idxfilePrefixIter) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *idxfilePrefixIter) bucketOffset(pos int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *idxfilePrefixIter) bucketCRC32(pos int) uint32 { _ = "STUB: not implemented"; return 0 }

func (i *idxfilePrefixIter) Close() error { _ = "STUB: not implemented"; return nil }

type Entry struct {
	Hash   plumbing.Hash
	CRC32  uint32
	Offset uint64
}

type idxfileEntryOffsetIter struct {
	entries entriesByOffset
	pos     int
}

func (i *idxfileEntryOffsetIter) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *idxfileEntryOffsetIter) Close() error { _ = "STUB: not implemented"; return nil }

type entriesByOffset []*Entry

func (o entriesByOffset) Len() int { _ = "STUB: not implemented"; return 0 }

func (o entriesByOffset) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (o entriesByOffset) Swap(i, j int) { _ = "STUB: not implemented"; return }
