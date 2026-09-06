package commitgraph

import (
	"github.com/emirpasic/gods/trees/binaryheap"

	"github.com/go-git/go-git/v6/plumbing"
)

type commitNodeStackable interface {
	Push(c CommitNode)
	Pop() (CommitNode, bool)
	Peek() (CommitNode, bool)
	Size() int
}

type commitNodeLifo struct {
	l []CommitNode
}

func (l *commitNodeLifo) Push(c CommitNode) { _ = "STUB: not implemented"; return }

func (l *commitNodeLifo) Pop() (CommitNode, bool) {
	_ = "STUB: not implemented"
	return *new(CommitNode), false
}

func (l *commitNodeLifo) Peek() (CommitNode, bool) {
	_ = "STUB: not implemented"
	return *new(CommitNode), false
}

func (l *commitNodeLifo) Size() int { _ = "STUB: not implemented"; return 0 }

type commitNodeHeap struct {
	*binaryheap.Heap
}

func (h *commitNodeHeap) Push(c CommitNode) { _ = "STUB: not implemented"; return }

func (h *commitNodeHeap) Pop() (CommitNode, bool) {
	_ = "STUB: not implemented"
	return *new(CommitNode), false
}

func (h *commitNodeHeap) Peek() (CommitNode, bool) {
	_ = "STUB: not implemented"
	return *new(CommitNode), false
}

func (h *commitNodeHeap) Size() int { _ = "STUB: not implemented"; return 0 }

func generationAndDateOrderComparator(left, right any) int { _ = "STUB: not implemented"; return 0 }

func composeIgnores(ignore []plumbing.Hash, seenExternal map[plumbing.Hash]bool) map[plumbing.Hash]struct{} {
	_ = "STUB: not implemented"
	return nil
}
