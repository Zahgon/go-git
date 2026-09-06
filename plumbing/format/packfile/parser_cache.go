package packfile

import (
	"github.com/go-git/go-git/v6/plumbing"
)

const maxObjectsPrealloc = 1 << 16

func newParserCache() *parserCache { _ = "STUB: not implemented"; return nil }

type parserCache struct {
	oi         []*ObjectHeader
	oiByHash   map[plumbing.Hash]*ObjectHeader
	oiByOffset map[int64]*ObjectHeader
}

func (c *parserCache) Add(oh *ObjectHeader) { _ = "STUB: not implemented"; return }

func (c *parserCache) Reset(n int) { _ = "STUB: not implemented"; return }
