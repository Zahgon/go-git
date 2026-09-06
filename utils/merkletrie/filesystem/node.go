package filesystem

import (
	iofs "io/fs"
	"os"
	"time"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/gitignore"
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

var ignore = map[string]bool{
	".git": true,
}

type Options struct {
	AutoCRLF bool

	Index *index.Index

	IgnoreScope *gitignore.Scope
}

type node struct {
	fs         billy.Filesystem
	submodules map[string]plumbing.Hash
	idx        *index.Index
	idxMap     map[string]*index.Entry

	trackedDirs map[string]struct{}

	options *Options

	scope         *gitignore.Scope
	scopeResolved bool

	path     string
	hash     []byte
	children []noder.Noder
	isDir    bool
	mode     os.FileMode
	size     int64
	modTime  time.Time
}

func NewRootNode(
	fs billy.Filesystem,
	submodules map[string]plumbing.Hash,
) noder.Noder {
	_ = "STUB: not implemented"
	return *new(noder.Noder)
}

func NewRootNodeWithOptions(
	fs billy.Filesystem,
	submodules map[string]plumbing.Hash,
	options Options,
) noder.Noder {
	_ = "STUB: not implemented"
	return *new(noder.Noder)
}

func (n *node) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (n *node) Name() string { _ = "STUB: not implemented"; return "" }

func (n *node) IsDir() bool { _ = "STUB: not implemented"; return false }

func (n *node) Skip() bool { _ = "STUB: not implemented"; return false }

func (n *node) Children() ([]noder.Noder, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *node) NumChildren() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *node) calculateChildren() error { _ = "STUB: not implemented"; return nil }

func (n *node) resolveScope(files []iofs.DirEntry) error { _ = "STUB: not implemented"; return nil }

func (n *node) pathComponents() []string { _ = "STUB: not implemented"; return nil }

func (n *node) shouldSkipIgnored(name string, isDir bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *node) newChildNode(file os.FileInfo) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *node) calculateHash() { _ = "STUB: not implemented"; return }

func (n *node) metadataMatches(entry *index.Entry) bool { _ = "STUB: not implemented"; return false }

func (n *node) doCalculateHashForRegular() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}

func (n *node) doCalculateHashForSymlink() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}

func (n *node) String() string { _ = "STUB: not implemented"; return "" }
