package commitgraph

import (
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	commitgraph "github.com/go-git/go-git/v6/plumbing/format/commitgraph"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type graphCommitNode struct {
	hash plumbing.Hash

	index uint32

	commitData *commitgraph.CommitData
	gci        *graphCommitNodeIndex
}

type graphCommitNodeIndex struct {
	commitGraph commitgraph.Index
	s           storer.EncodedObjectStorer
}

func NewGraphCommitNodeIndex(commitGraph commitgraph.Index, s storer.EncodedObjectStorer) CommitNodeIndex {
	_ = "STUB: not implemented"
	return *new(CommitNodeIndex)
}

func (gci *graphCommitNodeIndex) Get(hash plumbing.Hash) (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

func (c *graphCommitNode) ID() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (c *graphCommitNode) Tree() (*object.Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *graphCommitNode) CommitTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *graphCommitNode) NumParents() int { _ = "STUB: not implemented"; return 0 }

func (c *graphCommitNode) ParentNodes() CommitNodeIter {
	_ = "STUB: not implemented"
	return *new(CommitNodeIter)
}

func (c *graphCommitNode) ParentNode(i int) (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

func (c *graphCommitNode) ParentHashes() []plumbing.Hash { _ = "STUB: not implemented"; return nil }

func (c *graphCommitNode) Generation() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *graphCommitNode) GenerationV2() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *graphCommitNode) Commit() (*object.Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *graphCommitNode) String() string { _ = "STUB: not implemented"; return "" }
