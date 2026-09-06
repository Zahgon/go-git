package commitgraph

import (
	"github.com/go-git/go-git/v6/plumbing"
)

func NewCommitNodeIterDateOrder(c CommitNode,
	seenExternal map[plumbing.Hash]bool,
	ignore []plumbing.Hash,
) CommitNodeIter {
	_ = "STUB: not implemented"
	return *new(CommitNodeIter)
}
