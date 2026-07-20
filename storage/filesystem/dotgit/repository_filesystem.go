package dotgit

import (
	"io/fs"
	"os"

	"github.com/go-git/go-billy/v6"
)

type RepositoryFilesystem struct {
	dotGitFs       billy.Filesystem
	commonDotGitFs billy.Filesystem
}

func NewRepositoryFilesystem(dotGitFs, commonDotGitFs billy.Filesystem) *RepositoryFilesystem {
	_ = "STUB: not implemented"
	return nil
}

func (fs *RepositoryFilesystem) mapToRepositoryFsByPath(path string) billy.Filesystem {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem)
}

func (fs *RepositoryFilesystem) Create(filename string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (fs *RepositoryFilesystem) Open(filename string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (fs *RepositoryFilesystem) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (fs *RepositoryFilesystem) Stat(filename string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (fs *RepositoryFilesystem) Rename(oldpath, newpath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *RepositoryFilesystem) Remove(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *RepositoryFilesystem) Join(elem ...string) string { _ = "STUB: not implemented"; return "" }

func (fs *RepositoryFilesystem) TempFile(dir, prefix string) (billy.File, error) {
	_ = "STUB: not implemented"
	return *new(billy.File), nil
}

func (fs *RepositoryFilesystem) ReadDir(path string) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *RepositoryFilesystem) MkdirAll(filename string, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *RepositoryFilesystem) Lstat(filename string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

func (fs *RepositoryFilesystem) Symlink(target, link string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *RepositoryFilesystem) Readlink(link string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (fs *RepositoryFilesystem) Chroot(path string) (billy.Filesystem, error) {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem), nil
}

func (fs *RepositoryFilesystem) Root() string { _ = "STUB: not implemented"; return "" }
