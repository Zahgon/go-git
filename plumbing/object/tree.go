package object

import (
	"context"
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/filemode"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

const (
	maxTreeDepth      = 1024
	startingStackSize = 8
)

var (
	ErrMaxTreeDepth      = errors.New("maximum tree depth exceeded")
	ErrFileNotFound      = errors.New("file not found")
	ErrDirectoryNotFound = errors.New("directory not found")
	ErrEntryNotFound     = errors.New("entry not found")
	ErrEntriesNotSorted  = errors.New("entries in tree are not sorted")
	ErrMalformedTree     = errors.New("malformed tree")
	ErrDuplicateEntry    = errors.New("duplicate entry in tree")
	ErrInvalidTree       = errors.New("invalid tree")
)

const maxTreeEntryNameLen = 4096

type Tree struct {
	Entries []TreeEntry
	Hash    plumbing.Hash

	s             storer.EncodedObjectStorer
	t             map[string]*Tree
	entriesSorted bool
}

func GetTree(s storer.EncodedObjectStorer, h plumbing.Hash) (*Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodeTree(s storer.EncodedObjectStorer, o plumbing.EncodedObject) (*Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TreeEntry struct {
	Name string
	Mode filemode.FileMode
	Hash plumbing.Hash
}

func (t *Tree) File(path string) (*File, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) Size(path string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (t *Tree) Tree(path string) (*Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) TreeEntryFile(e *TreeEntry) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tree) FindEntry(path string) (*TreeEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tree) dir(baseName string) (*Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) entry(baseName string) (*TreeEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tree) searchEntry(baseName string) *TreeEntry { _ = "STUB: not implemented"; return nil }

func (t *Tree) searchEntryIndex(name string) int { _ = "STUB: not implemented"; return 0 }

func (t *Tree) Files() *FileIter { _ = "STUB: not implemented"; return nil }

func (t *Tree) ID() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (t *Tree) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (t *Tree) reset() { _ = "STUB: not implemented"; return }

func (t *Tree) Decode(o plumbing.EncodedObject) (err error) { _ = "STUB: not implemented"; return nil }

type TreeEntrySorter []TreeEntry

func (s TreeEntrySorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (s TreeEntrySorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s TreeEntrySorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func treeEntrySortName(e *TreeEntry) string { _ = "STUB: not implemented"; return "" }

func canonicalTreeMode(mode filemode.FileMode) filemode.FileMode {
	_ = "STUB: not implemented"
	return *new(filemode.FileMode)
}

func (t *Tree) Encode(o plumbing.EncodedObject) (err error) { _ = "STUB: not implemented"; return nil }

func (t *Tree) Validate() error { _ = "STUB: not implemented"; return nil }

func isValidTreeMode(mode filemode.FileMode) bool { _ = "STUB: not implemented"; return false }

func (t *Tree) Diff(to *Tree) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}

func (t *Tree) DiffContext(ctx context.Context, to *Tree) (Changes, error) {
	_ = "STUB: not implemented"
	return *new(Changes), nil
}

func (t *Tree) Patch(to *Tree) (*Patch, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tree) PatchContext(ctx context.Context, to *Tree) (*Patch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type treeEntryIter struct {
	t   *Tree
	pos int
}

func (iter *treeEntryIter) Next() (TreeEntry, error) {
	_ = "STUB: not implemented"
	return *new(TreeEntry), nil
}

type TreeWalker struct {
	stack     []*treeEntryIter
	base      string
	recursive bool
	seen      map[plumbing.Hash]bool

	skipPathValidation bool

	s storer.EncodedObjectStorer
	t *Tree
}

func NewTreeWalker(t *Tree, recursive bool, seen map[plumbing.Hash]bool) *TreeWalker {
	_ = "STUB: not implemented"
	return nil
}

func (w *TreeWalker) Next() (name string, entry TreeEntry, err error) {
	_ = "STUB: not implemented"
	return "", *new(TreeEntry), nil
}

func (w *TreeWalker) Tree() *Tree { _ = "STUB: not implemented"; return nil }

func (w *TreeWalker) Close() { _ = "STUB: not implemented"; return }

type TreeIter struct {
	storer.EncodedObjectIter
	s storer.EncodedObjectStorer
}

func NewTreeIter(s storer.EncodedObjectStorer, iter storer.EncodedObjectIter) *TreeIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *TreeIter) Next() (*Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (iter *TreeIter) ForEach(cb func(*Tree) error) error { _ = "STUB: not implemented"; return nil }

func simpleJoin(parent, child string) string { _ = "STUB: not implemented"; return "" }
