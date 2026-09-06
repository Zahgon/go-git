package cache

import (
	"container/list"
	"sync"
)

type BufferLRU struct {
	MaxSize FileSize

	actualSize FileSize
	ll         *list.List
	cache      map[int64]*list.Element
	mut        sync.Mutex
}

func NewBufferLRU(maxSize FileSize) *BufferLRU { _ = "STUB: not implemented"; return nil }

func NewBufferLRUDefault() *BufferLRU { _ = "STUB: not implemented"; return nil }

type buffer struct {
	Key   int64
	Slice []byte
}

func (c *BufferLRU) Put(key int64, slice []byte) { _ = "STUB: not implemented"; return }

func (c *BufferLRU) Get(key int64) ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func (c *BufferLRU) Clear() { _ = "STUB: not implemented"; return }
