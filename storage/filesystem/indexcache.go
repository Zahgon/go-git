package filesystem

import (
	"sync"
	"time"

	"github.com/go-git/go-git/v6/plumbing/format/index"
)

type IndexCache interface {
	Get(modTime time.Time, fileSize int64) *index.Index

	Set(idx *index.Index, modTime time.Time, fileSize int64)

	Clear()
}

type statIndexCache struct {
	mu       sync.RWMutex
	cached   *index.Index
	modTime  time.Time
	fileSize int64
}

func NewIndexCache() IndexCache { _ = "STUB: not implemented"; return *new(IndexCache) }

func (c *statIndexCache) Get(modTime time.Time, fileSize int64) *index.Index {
	_ = "STUB: not implemented"
	return nil
}

func (c *statIndexCache) Set(idx *index.Index, modTime time.Time, fileSize int64) {
	_ = "STUB: not implemented"
	return
}

func (c *statIndexCache) Clear() { _ = "STUB: not implemented"; return }
