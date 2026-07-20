package binary

import (
	"bufio"
	"errors"
	"io"
	"sync"
)

var ErrIntegerOverflow = errors.New("variable-width integer overflow")

func Read(r io.Reader, data ...any) error { _ = "STUB: not implemented"; return nil }

func ReadUntil(r io.Reader, delim byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ReadUntilFromBufioReader(r *bufio.Reader, delim byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadVariableWidthInt(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

const (
	maskContinue = uint8(128)
	maskLength   = uint8(127)
	lengthBits   = uint8(7)
)

func ReadUint64(r io.Reader) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadUint32(r io.Reader) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadUint16(r io.Reader) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

const sniffLen = 8000

var sniffPool = sync.Pool{
	New: func() any {
		b := make([]byte, sniffLen)
		return &b
	},
}

func IsBinary(r io.Reader) (bool, error) { _ = "STUB: not implemented"; return false, nil }
