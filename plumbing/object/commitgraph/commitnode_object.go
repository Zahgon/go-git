package commitgraph

import (
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type objectCommitNode struct {
	nodeIndex CommitNodeIndex
	commit    *object.Commit
}

func NewObjectCommitNodeIndex(s storer.EncodedObjectStorer) CommitNodeIndex {
	_ = "STUB: not implemented"
	return *new(CommitNodeIndex)
}

func (oci *objectCommitNodeIndex) Get(hash plumbing.Hash) (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

type objectCommitNodeIndex struct {
	s storer.EncodedObjectStorer
}

func (c *objectCommitNode) CommitTime() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *objectCommitNode) ID() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}

func (c *objectCommitNode) Tree() (*object.Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *objectCommitNode) NumParents() int { _ = "STUB: not implemented"; return 0 }

func (c *objectCommitNode) ParentNodes() CommitNodeIter {
	_ = "STUB: not implemented"
	return *new(CommitNodeIter)
}

func (c *objectCommitNode) ParentNode(i int) (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

func (c *objectCommitNode) ParentHashes() []plumbing.Hash { _ = "STUB: not implemented"; return nil }

func (c *objectCommitNode) Generation() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *objectCommitNode) GenerationV2() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *objectCommitNode) Commit() (*object.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
