package cache

import (
	"container/list"
	"sync"

	"github.com/go-git/go-git/v6/plumbing"
)

type ObjectLRU struct {
	MaxSize FileSize

	actualSize FileSize
	ll         *list.List
	cache      map[any]*list.Element
	mut        sync.Mutex
}

func NewObjectLRU(maxSize FileSize) *ObjectLRU { _ = "STUB: not implemented"; return nil }

func NewObjectLRUDefault() *ObjectLRU { _ = "STUB: not implemented"; return nil }

func (c *ObjectLRU) Put(obj plumbing.EncodedObject) { _ = "STUB: not implemented"; return }

func (c *ObjectLRU) Get(k plumbing.Hash) (plumbing.EncodedObject, bool) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), false
}

func (c *ObjectLRU) Clear() { _ = "STUB: not implemented"; return }
