package packfile

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing/storer"
)

var signature = []byte{'P', 'A', 'C', 'K'}

const (
	VersionSupported uint32 = 2

	firstLengthBits = uint8(4)
	lengthBits      = uint8(7)
	maskFirstLength = 15
	maskContinue    = 0x80
	maskLength      = uint8(127)
)

func UpdateObjectStorage(s storer.Storer, packfile io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func WritePackfileToObjectStorage(
	sw storer.PackfileWriter,
	packfile io.Reader,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ValidateOFSDeltaBase(deltaOffset, negativeOffset int64) error {
	_ = "STUB: not implemented"
	return nil
}
