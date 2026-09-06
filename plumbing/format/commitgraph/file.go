package commitgraph

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

var (
	ErrUnsupportedVersion = errors.New("unsupported version")

	ErrUnsupportedHash = errors.New("unsupported hash algorithm")

	ErrMalformedCommitGraphFile = errors.New("malformed commit graph file")

	ErrTooManyChunks = errors.New("commitgraph: too many chunks")

	ErrParentNotInIndex = errors.New("commitgraph: parent is not part of the index being encoded")

	commitFileSignature = []byte{'C', 'G', 'P', 'H'}

	parentNone        = uint32(0x70000000)
	parentOctopusUsed = uint32(0x80000000)
	parentOctopusMask = uint32(0x7fffffff)
	parentLast        = uint32(0x80000000)
)

const (
	szUint32 = 4
	szUint64 = 8

	szSignature  = 4
	szHeader     = 4
	szCommitData = 2*szUint32 + szUint64

	lenFanout = 256
)

type sizer interface {
	Size() int64
}

func readerSize(r io.ReaderAt) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

type fileIndex struct {
	reader                ReaderAtCloser
	fanout                [lenFanout]uint32
	offsets               [lenChunks]int64
	sizes                 [lenChunks]int64
	parent                Index
	hasGenerationV2       bool
	minimumNumberOfHashes uint32
	objSize               int
	numChunks             uint8
	fileSize              int64
}

type ReaderAtCloser interface {
	io.ReaderAt
	io.Closer
}

func OpenFileIndex(reader ReaderAtCloser) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func OpenFileIndexWithParent(reader ReaderAtCloser, parent Index) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func (fi *fileIndex) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (fi *fileIndex) verifyFileHeader() error { _ = "STUB: not implemented"; return nil }

func (fi *fileIndex) verifyFileSize() error { _ = "STUB: not implemented"; return nil }

type chunkAssignment struct {
	ct     ChunkType
	offset int64
}

func (fi *fileIndex) readChunkHeaders() error { _ = "STUB: not implemented"; return nil }

func (fi *fileIndex) verifyChunkSizes() error { _ = "STUB: not implemented"; return nil }

func (fi *fileIndex) readFanout() error { _ = "STUB: not implemented"; return nil }

func (fi *fileIndex) GetIndexByHash(h plumbing.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (fi *fileIndex) GetCommitDataByIndex(idx uint32) (*CommitData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fi *fileIndex) GetHashByIndex(idx uint32) (found plumbing.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (fi *fileIndex) getHashesFromIndexes(indexes []uint32) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fi *fileIndex) Hashes() []plumbing.Hash { _ = "STUB: not implemented"; return nil }

func (fi *fileIndex) HasGenerationV2() bool { _ = "STUB: not implemented"; return false }

func (fi *fileIndex) MaximumNumberOfHashes() uint32 { _ = "STUB: not implemented"; return 0 }
