package index

import (
	"errors"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/filemode"
)

var (
	ErrUnsupportedVersion = errors.New("unsupported version")

	ErrEntryNotFound = errors.New("entry not found")

	indexSignature              = []byte{'D', 'I', 'R', 'C'}
	treeExtSignature            = []byte{'T', 'R', 'E', 'E'}
	resolveUndoExtSignature     = []byte{'R', 'E', 'U', 'C'}
	endOfIndexEntryExtSignature = []byte{'E', 'O', 'I', 'E'}
)

type Stage int

const (
	Merged Stage = 1

	AncestorMode Stage = 1

	OurMode Stage = 2

	TheirMode Stage = 3
)

type Index struct {
	Version uint32

	Entries []*Entry

	Cache *Tree

	ResolveUndo *ResolveUndo

	EndOfIndexEntry *EndOfIndexEntry

	ModTime time.Time
}

func (i *Index) Add(path string) (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *Index) Entry(path string) (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *Index) Remove(path string) (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *Index) Glob(pattern string) (matches []*Entry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Index) String() string { _ = "STUB: not implemented"; return "" }

type Entry struct {
	Hash plumbing.Hash

	Name string

	CreatedAt time.Time

	ModifiedAt time.Time

	Dev, Inode uint32

	Mode filemode.FileMode

	UID, GID uint32

	Size uint32

	Stage Stage

	SkipWorktree bool

	IntentToAdd bool
}

func (e Entry) String() string { _ = "STUB: not implemented"; return "" }

type Tree struct {
	Entries []TreeEntry
}

type TreeEntry struct {
	Path string

	Entries int

	Trees int

	Hash plumbing.Hash
}

type ResolveUndo struct {
	Entries []ResolveUndoEntry
}

type ResolveUndoEntry struct {
	Path   string
	Stages map[Stage]plumbing.Hash
}

type EndOfIndexEntry struct {
	Offset uint32

	Hash plumbing.Hash
}

func (i *Index) SkipUnless(patterns []string) { _ = "STUB: not implemented"; return }
