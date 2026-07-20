//go:build darwin || linux

package mmap

import (
	"github.com/go-git/go-billy/v6"
)

var (
	idxSignature = []byte{255, 't', 'O', 'c'}
	idxMinLen    = idxHeaderSize + idxFanoutSize + idxCrcSize + len(idxSignature) + 40
	idxSupported = uint32(2)
)

const (
	idxHeaderSize = 8
	idxFanoutSize = 256 * 4
	idxCrcSize    = 4

	off32Size = 4
	off64Size = 8

	is64bitsMask = uint64(1) << 31
)

func (s *PackScanner) loadIdxFile(idx billy.File) error { _ = "STUB: not implemented"; return nil }
