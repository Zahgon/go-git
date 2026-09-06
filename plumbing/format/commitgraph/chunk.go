package commitgraph

const (
	szChunkSig     = 4
	chunkSigOffset = 4
)

var chunkSignatures = []byte("OIDFOIDLCDATGDA2GDO2EDGEBIDXBDATBASE\000\000\000\000")

type ChunkType int

const (
	OIDFanoutChunk ChunkType = iota
	OIDLookupChunk
	CommitDataChunk
	GenerationDataChunk
	GenerationDataOverflowChunk
	ExtraEdgeListChunk
	BloomFilterIndexChunk
	BloomFilterDataChunk
	BaseGraphsListChunk
	ZeroChunk
)
const lenChunks = int(ZeroChunk)

func (ct ChunkType) Signature() []byte { _ = "STUB: not implemented"; return nil }

func ChunkTypeFromBytes(b []byte) (ChunkType, bool) {
	_ = "STUB: not implemented"
	return *new(ChunkType), false
}
