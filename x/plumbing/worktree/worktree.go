package worktree

import (
	"errors"
	"regexp"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/storage"
	xstorage "github.com/go-git/go-git/v6/x/storage"
)

const (
	dotgitDir    = ".git"
	worktrees    = "worktrees"
	commonDir    = "commondir"
	gitDir       = "gitdir"
	head         = "HEAD"
	originalHead = "ORIG_HEAD"
	refs         = "refs"

	dirMode               = 0o777
	worktreeDotGitMaxSize = 1024
)

var (
	worktreeNameRE = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)

	ErrWorktreeNotFound = errors.New("worktree not found")

	ErrWorktreeAlreadyExists = errors.New("worktree already exists")
)

type Worktree struct {
	storer xstorage.WorktreeStorer
}

func New(storer storage.Storer) (*Worktree, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Worktree) Add(wt billy.Filesystem, name string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) List() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Worktree) Open(wt billy.Filesystem) (*git.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Worktree) Init(wt billy.Filesystem, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) getDualFS(wt billy.Filesystem) billy.Filesystem {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem)
}

func (w *Worktree) addDotGitDirs(wt billy.Filesystem, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) addWorktreeDotGitFile(wt billy.Filesystem, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) addDotGitFiles(dotgit, wt billy.Filesystem, name string, opts *options) error {
	_ = "STUB: not implemented"
	return nil
}

func writeFile(wt billy.Filesystem, fn string, data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func path(wtn, fn string) string { _ = "STUB: not implemented"; return "" }
