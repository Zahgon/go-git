package commitgraph

import (
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

type CommitNode interface {
	ID() plumbing.Hash

	Tree() (*object.Tree, error)

	CommitTime() time.Time

	NumParents() int

	ParentNodes() CommitNodeIter

	ParentNode(i int) (CommitNode, error)

	ParentHashes() []plumbing.Hash

	Generation() uint64

	GenerationV2() uint64

	Commit() (*object.Commit, error)
}

type CommitNodeIndex interface {
	Get(hash plumbing.Hash) (CommitNode, error)
}

type CommitNodeIter interface {
	Next() (CommitNode, error)
	ForEach(func(CommitNode) error) error
	Close()
}

type parentCommitNodeIter struct {
	node CommitNode
	i    int
}

func newParentgraphCommitNodeIter(node CommitNode) CommitNodeIter {
	_ = "STUB: not implemented"
	return *new(CommitNodeIter)
}

func (iter *parentCommitNodeIter) Next() (CommitNode, error) {
	_ = "STUB: not implemented"
	return *new(CommitNode), nil
}

func (iter *parentCommitNodeIter) ForEach(cb func(CommitNode) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (iter *parentCommitNodeIter) Close() { _ = "STUB: not implemented"; return }
