package diff

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/filemode"
)

type Operation int

const (
	Equal Operation = iota

	Add

	Delete
)

type Patch interface {
	FilePatches() []FilePatch

	Message() string
}

type FilePatch interface {
	IsBinary() bool

	Files() (from, to File)

	Chunks() []Chunk
}

type File interface {
	Hash() plumbing.Hash

	Mode() filemode.FileMode

	Path() string
}

type Chunk interface {
	Content() string

	Type() Operation
}
