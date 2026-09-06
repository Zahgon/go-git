package git

import (
	"errors"
	"io/fs"
	"os"

	"github.com/go-git/go-billy/v6"
)

func defaultProtectHFS() bool { _ = "STUB: not implemented"; return false }

func defaultProtectNTFS() bool { _ = "STUB: not implemented"; return false }

type worktreeFilesystem struct {
	billy.Filesystem
	protectNTFS bool
	protectHFS  bool
}

func newWorktreeFilesystem(fs billy.Filesystem, protectNTFS, protectHFS bool) *worktreeFilesystem {
	_ = "STUB: not implemented"
	return nil
}

func (sfs *worktreeFilesystem) Create(filename string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (sfs *worktreeFilesystem) Open(filename string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (sfs *worktreeFilesystem) OpenFile(filename string, flag int, perm fs.FileMode) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (sfs *worktreeFilesystem) Stat(filename string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (sfs *worktreeFilesystem) Remove(filename string) error { _ = "STUB: not implemented"; return nil }

func (sfs *worktreeFilesystem) Rename(from, to string) error { _ = "STUB: not implemented"; return nil }

func (sfs *worktreeFilesystem) ReadDir(path string) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sfs *worktreeFilesystem) Lstat(filename string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (sfs *worktreeFilesystem) Symlink(target, link string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sfs *worktreeFilesystem) Readlink(link string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (sfs *worktreeFilesystem) MkdirAll(path string, perm fs.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (sfs *worktreeFilesystem) TempFile(_, _ string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (sfs *worktreeFilesystem) validReadPath(p string) error { _ = "STUB: not implemented"; return nil }

func (sfs *worktreeFilesystem) Chroot(path string) (billy.Filesystem, error) {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem), nil
}

var errUnsupportedOperation = errors.New("unsupported operation")

func isDotGitVariant(part string, protectHFS bool) bool { _ = "STUB: not implemented"; return false }

func (sfs *worktreeFilesystem) validPath(paths ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sfs *worktreeFilesystem) validWritePath(paths ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sfs *worktreeFilesystem) validNoLeadingSymlink(paths ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sfs *worktreeFilesystem) validSymlinkName(name string) error {
	_ = "STUB: not implemented"
	return nil
}
