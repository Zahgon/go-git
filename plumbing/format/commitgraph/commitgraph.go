package commitgraph

import (
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
)

type CommitData struct {
	TreeHash plumbing.Hash

	ParentIndexes []uint32

	ParentHashes []plumbing.Hash

	Generation uint64

	GenerationV2 uint64

	When time.Time
}

func (c *CommitData) GenerationV2Data() uint64 { _ = "STUB: not implemented"; return 0 }

type Index interface {
	GetIndexByHash(h plumbing.Hash) (uint32, error)

	GetHashByIndex(i uint32) (plumbing.Hash, error)

	GetCommitDataByIndex(i uint32) (*CommitData, error)

	Hashes() []plumbing.Hash

	HasGenerationV2() bool

	MaximumNumberOfHashes() uint32

	io.Closer
}
