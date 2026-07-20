package fdpool

import (
	"container/list"
	"sync"
)

type Member interface {
	ReleaseNow() error
}

type Pinnable interface {
	Pinned() bool
}

type Handle struct {
	elem *list.Element
}

type entry struct {
	m Member
	h *Handle
}

type Stats struct {
	Capacity int

	Active int

	Hits uint64

	Evictions uint64

	EvictionFailures uint64

	PinnedSkips uint64
}

type Pool struct {
	mu               sync.Mutex
	capacity         int
	lru              *list.List
	hits             uint64
	evictions        uint64
	evictionFailures uint64
	pinnedSkips      uint64
}

func New(capacity int) *Pool { _ = "STUB: not implemented"; return nil }

func (p *Pool) Touch(m Member, h *Handle) { _ = "STUB: not implemented"; return }

func (p *Pool) Forget(h *Handle) { _ = "STUB: not implemented"; return }

func (p *Pool) Stats() Stats { _ = "STUB: not implemented"; return *new(Stats) }
