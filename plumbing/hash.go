package plumbing

import (
	"hash"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

type Hash = ObjectID

var ZeroHash ObjectID

func NewHash(s string) Hash { _ = "STUB: not implemented"; return *new(Hash) }

type Hasher struct {
	hash.Hash
	format format.ObjectFormat
}

func NewHasher(f format.ObjectFormat, t ObjectType, size int64) Hasher {
	_ = "STUB: not implemented"
	return *new(Hasher)
}

func (h Hasher) Reset(t ObjectType, size int64) { _ = "STUB: not implemented"; return }

func (h Hasher) Sum() (hash Hash) { _ = "STUB: not implemented"; return *new(Hash) }

func HashesSort(a []Hash) { _ = "STUB: not implemented"; return }

type HashSlice []Hash

func (p HashSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (p HashSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (p HashSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func IsHash(s string) bool { _ = "STUB: not implemented"; return false }
