package object

import (
	"github.com/go-git/go-git/v6/plumbing/filemode"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type File struct {
	Name string

	Mode filemode.FileMode

	Blob
}

func NewFile(name string, m filemode.FileMode, b *Blob) *File {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) Contents() (content string, err error) { _ = "STUB: not implemented"; return "", nil }

func (f *File) IsBinary() (bin bool, err error) { _ = "STUB: not implemented"; return false, nil }

func (f *File) Lines() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type FileIter struct {
	s storer.EncodedObjectStorer
	w TreeWalker
}

func NewFileIter(s storer.EncodedObjectStorer, t *Tree) *FileIter {
	_ = "STUB: not implemented"
	return nil
}

func (iter *FileIter) Next() (*File, error) { _ = "STUB: not implemented"; return nil, nil }

func (iter *FileIter) ForEach(cb func(*File) error) error { _ = "STUB: not implemented"; return nil }

func (iter *FileIter) Close() { _ = "STUB: not implemented"; return }
