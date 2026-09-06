package filesystem

import (
	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/cache"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
	"github.com/go-git/go-git/v6/x/fdpool"
)

const defaultPoolCapacity = 256

type Storage struct {
	fs     billy.Filesystem
	dir    *dotgit.DotGit
	hasher plumbing.Hasher

	*ObjectStorage
	ReferenceStorage
	IndexStorage
	ShallowStorage
	ConfigStorage
	ModuleStorage
	ReflogStorage
}

var (
	_ storer.IdleReleaser = (*Storage)(nil)
	_ storer.IdleReleaser = (*dotgit.DotGit)(nil)
)

type Options struct {
	ExclusiveAccess bool

	LargeObjectThreshold int64

	AlternatesFS billy.Filesystem

	HighMemoryMode bool

	ObjectFormat formatcfg.ObjectFormat

	UseInMemoryIdx bool

	IndexCache IndexCache

	Pool *fdpool.Pool
}

func NewStorage(fs billy.Filesystem, cache cache.Object) *Storage {
	_ = "STUB: not implemented"
	return nil
}

func NewStorageWithOptions(fs billy.Filesystem, c cache.Object, ops Options) *Storage {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) SetObjectFormat(of formatcfg.ObjectFormat) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) SupportsExtension(name, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Storage) Filesystem() billy.Filesystem {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem)
}

func (s *Storage) Init() error { _ = "STUB: not implemented"; return nil }

func (s *Storage) AddAlternate(remote string) error { _ = "STUB: not implemented"; return nil }

func (s *Storage) LowMemoryMode() bool { _ = "STUB: not implemented"; return false }
