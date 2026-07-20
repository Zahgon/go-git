package transport

import (
	"net/url"

	"github.com/go-git/go-billy/v6"
	"github.com/go-git/go-billy/v6/osfs"

	"github.com/go-git/go-git/v6/storage"
)

type Loader interface {
	Load(u *url.URL) (storage.Storer, error)
}

var DefaultLoader Loader = NewFilesystemLoader(osfs.New(""), false)

type FilesystemLoader struct {
	base   billy.Filesystem
	strict bool
}

func NewFilesystemLoader(base billy.Filesystem, strict bool) *FilesystemLoader {
	_ = "STUB: not implemented"
	return nil
}

func (l *FilesystemLoader) Load(u *url.URL) (storage.Storer, error) {
	_ = "STUB: not implemented"
	return *new(storage.Storer), nil
}

func (l *FilesystemLoader) load(path string, tried bool) (storage.Storer, error) {
	_ = "STUB: not implemented"
	return *new(storage.Storer), nil
}

func readGitfile(fs billy.Filesystem) (string, error) { _ = "STUB: not implemented"; return "", nil }

type MapLoader map[string]storage.Storer

func (l MapLoader) Load(u *url.URL) (storage.Storer, error) {
	_ = "STUB: not implemented"
	return *new(storage.Storer), nil
}
