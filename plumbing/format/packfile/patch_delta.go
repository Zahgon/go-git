package packfile

import (
	"bytes"
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

var (
	ErrInvalidDelta = errors.New("invalid delta")
	ErrDeltaCmd     = errors.New("wrong delta command")
)

const (
	maxPatchPreemptionSize uint = 65536

	minDeltaSize = 2
)

type offset struct {
	mask  byte
	shift uint
}

var offsets = []offset{
	{mask: 0x01, shift: 0},
	{mask: 0x02, shift: 8},
	{mask: 0x04, shift: 16},
	{mask: 0x08, shift: 24},
}

var sizes = []offset{
	{mask: 0x10, shift: 0},
	{mask: 0x20, shift: 8},
	{mask: 0x40, shift: 16},
}

func ApplyDelta(target, base plumbing.EncodedObject, delta *bytes.Buffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func PatchDelta(src, delta []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ReaderFromDelta(base plumbing.EncodedObject, deltaRC io.Reader) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func patchDelta(dst *bytes.Buffer, src, delta []byte) error { _ = "STUB: not implemented"; return nil }

func patchDeltaWriter(dst io.Writer, base io.ReaderAt, delta io.Reader,
	typ plumbing.ObjectType, writeHeader objectHeaderWriter, of format.ObjectFormat,
) (uint, plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return 0, *new(plumbing.Hash), nil
}

func isCopyFromSrc(cmd byte) bool { _ = "STUB: not implemented"; return false }

func isCopyFromDelta(cmd byte) bool { _ = "STUB: not implemented"; return false }

func decodeOffsetByteReader(cmd byte, delta io.ByteReader) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeOffset(cmd byte, delta []byte) (uint, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func decodeSizeByteReader(cmd byte, delta io.ByteReader) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeSize(cmd byte, delta []byte) (uint, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func invalidSize(sz, remaining uint) bool { _ = "STUB: not implemented"; return false }

func invalidOffsetSize(offset, sz, srcSz uint) bool { _ = "STUB: not implemented"; return false }

func sumOverflows(a, b uint) bool { _ = "STUB: not implemented"; return false }
