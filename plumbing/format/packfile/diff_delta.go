package packfile

import (
	"bytes"

	"github.com/go-git/go-git/v6/plumbing"
)

const (
	s = 16

	maxCopySize = 64 * 1024
)

func GetDelta(base, target plumbing.EncodedObject) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func getDelta(index *deltaIndex, base, target plumbing.EncodedObject) (o plumbing.EncodedObject, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func DiffDelta(src, tgt []byte) []byte { _ = "STUB: not implemented"; return nil }

func diffDelta(index *deltaIndex, src, tgt []byte) []byte { _ = "STUB: not implemented"; return nil }

func encodeInsertOperation(ibuf, buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func encodeCopyOperation(offset, length int) []byte { _ = "STUB: not implemented"; return nil }
