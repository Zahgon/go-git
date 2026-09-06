package commitgraph

import (
	"github.com/go-git/go-git/v6/plumbing"
)

type commitNodeIteratorTopological struct {
	exploreStack commitNodeStackable
	visitStack   commitNodeStackable
	inCounts     map[plumbing.Hash]int

	ignore map[plumbing.Hash]struct{}
}

func NewCommitNodeIterTopoOrder(c CommitNode,
	seenExternal map[plumbing.Hash]bool,
	ignore []plumbing.Hash,
) CommitNodeIter {
	_ = "STUB: not implemented"
	return *new(CommitNodeIter)
}

func (iter *commitNodeIteratorTopological) Next() (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

func (iter *commitNodeIteratorTopological) ForEach(cb func(CommitNode) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *commitNodeIteratorTopological) Close() { _ = "STUB: not implemented"; return }
