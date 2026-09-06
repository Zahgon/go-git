package index

import (
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

type node struct {
	path     string
	entry    *index.Entry
	children []noder.Noder
	isDir    bool
	skip     bool

	upholdExecutableBit bool
}

type RootNodeOptions struct {
	UpholdExecutableBit bool
}

func NewRootNode(idx *index.Index) noder.Noder { _ = "STUB: not implemented"; return *new(noder.Noder) }

func NewRootNodeWithOptions(idx *index.Index, options RootNodeOptions) noder.Noder {
	_ = "STUB: not implemented"
	return *new(noder.Noder)
}

func (n *node) String() string { _ = "STUB: not implemented"; return "" }

func (n *node) Skip() bool { _ = "STUB: not implemented"; return false }

func (n *node) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (n *node) Name() string { _ = "STUB: not implemented"; return "" }

func (n *node) IsDir() bool { _ = "STUB: not implemented"; return false }

func (n *node) Children() ([]noder.Noder, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *node) NumChildren() (int, error) { _ = "STUB: not implemented"; return 0, nil }
