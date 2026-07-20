package idxfile

import (
	"errors"
	"io"
	"io/fs"

	"github.com/go-git/go-git/v6/plumbing/hash"
)

var (
	ErrUnsupportedVersion = errors.New("unsupported version")

	ErrMalformedIdxFile = errors.New("malformed idx file")
)

const (
	fanout = 256
)

const (
	headerLen     = 8
	fanoutLen     = fanout * 4
	crc32Len      = 4
	offset32Len   = 4
	offset64Len   = 8
	trailerHashes = 2
)

type Input interface {
	io.Reader
	Stat() (fs.FileInfo, error)
}

type Decoder struct {
	in Input
	h  hash.Hash
}

func NewDecoder(in Input, h hash.Hash) *Decoder { _ = "STUB: not implemented"; return nil }

func (d *Decoder) Decode(idx *MemoryIndex) error { _ = "STUB: not implemented"; return nil }

func validateHeader(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readVersion(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readFanout(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readObjectNames(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readCRC32(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readOffsets(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readPackChecksum(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func readIdxChecksum(idx *MemoryIndex, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func validateIdxV2Size(idx *MemoryIndex, idxSize int64) error {
	_ = "STUB: not implemented"
	return nil
}

func minIdxV2Size(nr, hashsz int64) int64 { _ = "STUB: not implemented"; return 0 }

func maxIdxV2Size(nr, hashsz int64) int64 { _ = "STUB: not implemented"; return 0 }

func mulInt64(a, b int64) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func addInt64(a, b int64) (int64, bool) { _ = "STUB: not implemented"; return 0, false }
