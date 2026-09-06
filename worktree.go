package git

import (
	"context"
	"errors"
	"io/fs"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/gitignore"
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/utils/merkletrie"
)

var (
	ErrWorktreeNotClean = errors.New("worktree is not clean")

	ErrSubmoduleNotFound = errors.New("submodule not found")

	ErrUnstagedChanges = errors.New("worktree contains unstaged changes")

	ErrLocalChanges = errors.New("worktree contains local changes that would be overwritten by reset")

	ErrGitModulesSymlink = errors.New(gitmodulesFile + " is a symlink")

	ErrNonFastForwardUpdate = errors.New("non-fast-forward update")

	ErrRestoreWorktreeOnlyNotSupported = errors.New("worktree only is not supported")

	ErrSparseResetDirectoryNotFound = errors.New("sparse-reset directory not found on commit")
)

type Worktree struct {
	Excludes []gitignore.Pattern

	r          *Repository
	filesystem *worktreeFilesystem
}

func (w *Worktree) Filesystem() billy.Filesystem {
	_ = "STUB: not implemented"
	return *new(billy.Filesystem)
}

func (w *Worktree) reusableRootFS() (*worktreeFilesystem, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Worktree) Pull(o *PullOptions) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) PullContext(ctx context.Context, o *PullOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) updateSubmodules(ctx context.Context, o *SubmoduleUpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) Checkout(opts *CheckoutOptions) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) createBranch(opts *CheckoutOptions) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) getCommitFromCheckoutOptions(opts *CheckoutOptions) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) setHEADToCommit(commit plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) setHEADToBranch(branch plumbing.ReferenceName, commit plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) Reset(opts *ResetOptions) error { _ = "STUB: not implemented"; return nil }

func treeContainsDirs(tree *object.Tree, dirs []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *Worktree) Restore(o *RestoreOptions) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) resetIndex(t *object.Tree, dirs, files []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inFiles(files map[string]struct{}, v string) bool { _ = "STUB: not implemented"; return false }

func (w *Worktree) headTree() (*object.Tree, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Worktree) checkKeepResetConflicts(fromTree, toTree *object.Tree, sparseDirs, files []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) resetWorktreeToTree(cfg *config.Config, fromTree, toTree *object.Tree, files []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) resetWorktree(cfg *config.Config, t *object.Tree, files []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) checkoutChange(cfg *config.Config, fs *worktreeFilesystem, ch merkletrie.Change, t *object.Tree, idx *indexBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) containsUnstagedChanges(cfg *config.Config) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (w *Worktree) setHEADCommit(commit plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) checkoutChangeSubmodule(fs *worktreeFilesystem,
	name string,
	a merkletrie.Action,
	e *object.TreeEntry,
	idx *indexBuilder,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) checkoutChangeRegularFile(cfg *config.Config,
	fs *worktreeFilesystem,
	name string,
	a merkletrie.Action,
	t *object.Tree,
	e *object.TreeEntry,
	idx *indexBuilder,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) clearBlockingSymlinks(fs *worktreeFilesystem, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) checkoutFile(cfg *config.Config, fs *worktreeFilesystem, f *object.File) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) copyObjectToWorktree(cfg *config.Config, object *object.File, file billy.File) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) checkoutFileSymlink(fs *worktreeFilesystem, f *object.File) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) addIndexFromTreeEntry(name string, f *object.TreeEntry, idx *indexBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) addIndexFromFile(fs *worktreeFilesystem, name string, h plumbing.Hash, idx *indexBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) getTreeFromCommitHash(commit plumbing.Hash) (*object.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var fillSystemInfo func(e *index.Entry, sys any)

const gitmodulesFile = ".gitmodules"

func (w *Worktree) Submodule(name string) (*Submodule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveModuleURL(originURL, moduleURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (w *Worktree) Submodules() (Submodules, error) {
	_ = "STUB: not implemented"
	return *new(Submodules), nil
}

func (w *Worktree) submodulesWithConfig(cfg *config.Config) (Submodules, error) {
	_ = "STUB: not implemented"
	return *new(Submodules), nil
}

func (w *Worktree) newSubmodule(fromModules, fromConfig *config.Submodule) *Submodule {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) isSymlink(path string) bool { _ = "STUB: not implemented"; return false }

func (w *Worktree) readGitmodulesFile() (*config.Modules, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Worktree) Clean(opts *CleanOptions) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) doClean(status Status, opts *CleanOptions, dir string, files []fs.DirEntry) error {
	_ = "STUB: not implemented"
	return nil
}

type GrepResult struct {
	FileName string

	LineNumber int

	Content string

	TreeName string
}

func (gr GrepResult) String() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) Grep(opts *GrepOptions) ([]GrepResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Worktree) Grep(opts *GrepOptions) ([]GrepResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findMatchInFiles(fileiter *object.FileIter, treeName string, opts *GrepOptions) ([]GrepResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findMatchInFile(file *object.File, treeName string, opts *GrepOptions) ([]GrepResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rmFileAndDirsIfEmpty(fs billy.Filesystem, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func removeDirIfEmpty(fs billy.Filesystem, dir string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type indexBuilder struct {
	entries map[string]*index.Entry
}

func newIndexBuilder(idx *index.Index) *indexBuilder { _ = "STUB: not implemented"; return nil }

func (b *indexBuilder) Write(idx *index.Index) { _ = "STUB: not implemented"; return }

func (b *indexBuilder) Add(e *index.Entry) { _ = "STUB: not implemented"; return }

func (b *indexBuilder) Remove(name string) { _ = "STUB: not implemented"; return }

func buildFilePathMap(files []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }
