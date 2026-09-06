package packfile

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type Observer interface {
	OnHeader(count uint32) error

	OnInflatedObjectHeader(t plumbing.ObjectType, objSize, pos int64) error

	OnInflatedObjectContent(h plumbing.Hash, pos int64, crc uint32, content []byte) error

	OnFooter(h plumbing.Hash) error
}

type objectHeaderWriter func(typ plumbing.ObjectType, sz int64) error
