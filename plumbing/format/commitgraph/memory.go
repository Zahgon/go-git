package commitgraph

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type MemoryIndex struct {
	commitData      []commitData
	indexMap        map[plumbing.Hash]uint32
	hasGenerationV2 bool
}

type commitData struct {
	Hash plumbing.Hash
	*CommitData
}

func NewMemoryIndex() *MemoryIndex { _ = "STUB: not implemented"; return nil }

func (mi *MemoryIndex) GetIndexByHash(h plumbing.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mi *MemoryIndex) GetHashByIndex(i uint32) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (mi *MemoryIndex) GetCommitDataByIndex(i uint32) (*CommitData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mi *MemoryIndex) Hashes() []plumbing.Hash { _ = "STUB: not implemented"; return nil }

func (mi *MemoryIndex) Add(hash plumbing.Hash, data *CommitData) { _ = "STUB: not implemented"; return }

func (mi *MemoryIndex) HasGenerationV2() bool { _ = "STUB: not implemented"; return false }

func (mi *MemoryIndex) Close() error { _ = "STUB: not implemented"; return nil }

func (mi *MemoryIndex) MaximumNumberOfHashes() uint32 { _ = "STUB: not implemented"; return 0 }
