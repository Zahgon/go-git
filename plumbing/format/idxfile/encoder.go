package idxfile

import (
	"hash"
	"io"
)

type encoder struct {
	writer  io.Writer
	hashSum func() []byte
	idx     *MemoryIndex
}

type stateFnEncode func(*encoder) (stateFnEncode, error)

func Encode(w io.Writer, h hash.Hash, idx *MemoryIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func writeHeader(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeFanout(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeHashes(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeCRC32(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeOffsets(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeChecksums(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}
