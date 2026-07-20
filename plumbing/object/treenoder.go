package object

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/filemode"
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

type treeNoder struct {
	parent   *Tree
	name     string
	mode     filemode.FileMode
	hash     plumbing.Hash
	children []noder.Noder
}

func NewTreeRootNode(t *Tree) noder.Noder { _ = "STUB: not implemented"; return *new(noder.Noder) }

func (t *treeNoder) Skip() bool { _ = "STUB: not implemented"; return false }

func (t *treeNoder) isRoot() bool { _ = "STUB: not implemented"; return false }

func (t *treeNoder) String() string { _ = "STUB: not implemented"; return "" }

func (t *treeNoder) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (t *treeNoder) Name() string { _ = "STUB: not implemented"; return "" }

func (t *treeNoder) IsDir() bool { _ = "STUB: not implemented"; return false }

func (t *treeNoder) Children() ([]noder.Noder, error) { _ = "STUB: not implemented"; return nil, nil }

func transformChildren(t *Tree) ([]noder.Noder, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *treeNoder) NumChildren() (int, error) { _ = "STUB: not implemented"; return 0, nil }
