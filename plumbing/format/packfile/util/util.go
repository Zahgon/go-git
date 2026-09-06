package util

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

const (
	firstLengthBits = uint8(4)
	maskPayload     = 0x7f
	maskContinue    = 0x80
	maskType        = uint8(112)
)

var ErrLengthOverflow = errors.New("variable-length integer overflow")

func VariableLengthSize(first byte, reader io.ByteReader) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ObjectType(b byte) plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func EncodeLEB128(num uint) []byte { _ = "STUB: not implemented"; return nil }

func EncodeLEB128ToWriter(writer io.Writer, num uint) error { _ = "STUB: not implemented"; return nil }

func DecodeLEB128(input []byte) (uint, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func DecodeLEB128FromReader(input io.ByteReader) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

const uintBits = 32 << (^uint(0) >> 63)
