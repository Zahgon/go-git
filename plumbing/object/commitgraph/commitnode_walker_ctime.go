package commitgraph

import (
	"github.com/emirpasic/gods/trees/binaryheap"

	"github.com/go-git/go-git/v6/plumbing"
)

type commitNodeIteratorByCTime struct {
	heap         *binaryheap.Heap
	seenExternal map[plumbing.Hash]bool
	seen         map[plumbing.Hash]bool
}

func NewCommitNodeIterCTime(
	c CommitNode,
	seenExternal map[plumbing.Hash]bool,
	ignore []plumbing.Hash,
) CommitNodeIter {
	_ = "STUB: not implemented"
	return *new(CommitNodeIter)
}

func (w *commitNodeIteratorByCTime) Next() (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

func (w *commitNodeIteratorByCTime) ForEach(cb func(CommitNode) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *commitNodeIteratorByCTime) Close() { _ = "STUB: not implemented"; return }
