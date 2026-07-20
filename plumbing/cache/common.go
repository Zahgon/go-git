package cache

import "github.com/go-git/go-git/v6/plumbing"

const (
	Byte FileSize = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
)

type FileSize int64

const DefaultMaxSize FileSize = 96 * MiByte

type Object interface {
	Put(o plumbing.EncodedObject)

	Get(k plumbing.Hash) (plumbing.EncodedObject, bool)

	Clear()
}

type Buffer interface {
	Put(key int64, slice []byte)

	Get(key int64) ([]byte, bool)

	Clear()
}
