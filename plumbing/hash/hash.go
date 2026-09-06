package hash

import (
	"crypto"
	"errors"
	"hash"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

var ErrUnsupportedHashFunction = errors.New("unsupported hash function")

var algos = map[crypto.Hash]func() hash.Hash{}

func init() {
	reset()
}

func reset() { _ = "STUB: not implemented"; return }

func RegisterHash(h crypto.Hash, f func() hash.Hash) error { _ = "STUB: not implemented"; return nil }

type Hash interface {
	hash.Hash
}

func New(h crypto.Hash) Hash { _ = "STUB: not implemented"; return *new(Hash) }

func FromObjectFormat(f format.ObjectFormat) (hash.Hash, error) {
	_ = "STUB: not implemented"
	return *new(hash.Hash), nil
}
