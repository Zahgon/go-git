package revfile

import (
	"crypto"
	"errors"
	"hash"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

var (
	ErrUnsupportedVersion = errors.New("unsupported version")

	ErrMalformedRevFile = errors.New("malformed rev file")

	ErrUnsupportedHashFunction = errors.New("unsupported hash function")

	ErrEmptyReverseIndex = errors.New("reverse index is empty")

	revHeader = []byte{'R', 'I', 'D', 'X'}
)

const (
	VersionSupported        = 1
	sha1Hash         uint32 = 1
	sha256Hash       uint32 = 2
)

type decoder struct {
	reader  io.Reader
	hasher  crypto.Hash
	hash    hash.Hash
	version uint32

	objCount     int64
	packChecksum plumbing.ObjectID
	out          chan<- uint32
}

type stateFn func(*decoder) (stateFn, error)

func Decode(r io.Reader, objCount int64, packChecksum plumbing.ObjectID, out chan<- uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func readMagicNumber(d *decoder) (stateFn, error) {
	_ = "STUB: not implemented"
	return *new(stateFn), nil
}

func readVersion(d *decoder) (stateFn, error) { _ = "STUB: not implemented"; return *new(stateFn), nil }

func readHashFunction(d *decoder) (stateFn, error) {
	_ = "STUB: not implemented"
	return *new(stateFn), nil
}

func readEntries(d *decoder) (stateFn, error) { _ = "STUB: not implemented"; return *new(stateFn), nil }

func readPackChecksum(d *decoder) (stateFn, error) {
	_ = "STUB: not implemented"
	return *new(stateFn), nil
}

func readRevChecksum(d *decoder) (stateFn, error) {
	_ = "STUB: not implemented"
	return *new(stateFn), nil
}
