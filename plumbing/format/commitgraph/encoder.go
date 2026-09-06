package commitgraph

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/hash"
)

type Encoder struct {
	io.Writer
	hash hash.Hash
}

func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

func (e *Encoder) Encode(idx Index) error { _ = "STUB: not implemented"; return nil }

func lookupParentIndex(hashToIndex map[plumbing.Hash]uint32, h plumbing.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *Encoder) prepare(idx Index, hashes []plumbing.Hash) (hashToIndex map[plumbing.Hash]uint32, fanout []uint32, extraEdgesCount, generationV2OverflowCount uint32, err error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, 0, nil
}

func (e *Encoder) encodeFileHeader(chunkCount int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeChunkHeaders(chunkSignatures [][]byte, chunkSizes []uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeFanout(fanout []uint32) (err error) { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeOidLookup(hashes []plumbing.Hash) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeCommitData(hashes []plumbing.Hash, hashToIndex map[plumbing.Hash]uint32, idx Index) (extraEdges []uint32, generationV2Data []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (e *Encoder) encodeExtraEdges(extraEdges []uint32) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeGenerationV2Data(generationV2Data []uint64) (overflows []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Encoder) encodeGenerationV2Overflow(overflows []uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeChecksum() error { _ = "STUB: not implemented"; return nil }
