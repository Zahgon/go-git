package dotgit

import (
	"errors"
	"io"
	"io/fs"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/internal/packhandle"
	"github.com/go-git/go-git/v6/plumbing"
)

var errReadOnlyPack = errors.New("dotgit: pack file is read-only")

func (d *DotGit) OpenPackForReading(hash plumbing.Hash) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

type readOnlyPackFile struct {
	cursor packhandle.PackReader
	ra     io.ReaderAt
	name   string
	fs     billy.Filesystem
}

func (f *readOnlyPackFile) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (f *readOnlyPackFile) Close() error               { _ = "STUB: not implemented"; return nil }
func (f *readOnlyPackFile) Name() string               { _ = "STUB: not implemented"; return "" }
func (f *readOnlyPackFile) Seek(o int64, w int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *readOnlyPackFile) ReadAt(p []byte, off int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *readOnlyPackFile) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (f *readOnlyPackFile) Write(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (f *readOnlyPackFile) WriteAt(_ []byte, _ int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *readOnlyPackFile) Lock() error          { _ = "STUB: not implemented"; return nil }
func (f *readOnlyPackFile) Unlock() error        { _ = "STUB: not implemented"; return nil }
func (f *readOnlyPackFile) Truncate(int64) error { _ = "STUB: not implemented"; return nil }

var _ billy.File = (*readOnlyPackFile)(nil)
