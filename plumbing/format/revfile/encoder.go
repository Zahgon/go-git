package revfile

import (
	"hash"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
)

type encoder struct {
	writer io.Writer
	hash   hash.Hash

	entries      []uint32
	packChecksum plumbing.Hash
}

type stateFnEncode func(*encoder) (stateFnEncode, error)

func Encode(w io.Writer, h hash.Hash, idx *idxfile.MemoryIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *encoder) buildReverseIndex(idx *idxfile.MemoryIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func writeHeader(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeVersion(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeHashFunction(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeEntries(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writePackChecksum(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}

func writeRevChecksum(e *encoder) (stateFnEncode, error) {
	_ = "STUB: not implemented"
	return *new(stateFnEncode), nil
}
