package git

import (
	"errors"
	"io"
	"os"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/gitignore"
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/utils/merkletrie"
	"github.com/go-git/go-git/v6/utils/merkletrie/noder"
)

var (
	ErrDestinationExists = errors.New("destination exists")

	ErrGlobNoMatches = errors.New("glob pattern did not match any files")

	ErrUnsupportedStatusStrategy = errors.New("unsupported status strategy")
)

func (w *Worktree) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

type StatusOptions struct {
	Strategy StatusStrategy
}

func (w *Worktree) StatusWithOptions(o StatusOptions) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func (w *Worktree) status(cfg *config.Config, ss StatusStrategy, commit plumbing.Hash) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func nameFromAction(ch *merkletrie.Change) string { _ = "STUB: not implemented"; return "" }

func (w *Worktree) diffStagingWithWorktree(cfg *config.Config, reverse, excludeIgnoredChanges bool) (merkletrie.Changes, error) {
	_ = "STUB: not implemented"
	return *new(merkletrie.Changes), nil
}

func (w *Worktree) collectIgnorePatterns() []gitignore.Pattern {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) getSubmodulesStatus(cfg *config.Config) (map[string]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Worktree) diffCommitWithStaging(commit plumbing.Hash, reverse bool) (merkletrie.Changes, error) {
	_ = "STUB: not implemented"
	return *new(merkletrie.Changes), nil
}

func (w *Worktree) diffTreeWithStaging(t *object.Tree, reverse bool) (merkletrie.Changes, error) {
	_ = "STUB: not implemented"
	return *new(merkletrie.Changes), nil
}

func diffTrees(from, to *object.Tree) (merkletrie.Changes, error) {
	_ = "STUB: not implemented"
	return *new(merkletrie.Changes), nil
}

var emptyNoderHash = make([]byte, 24)

func diffTreeIsEquals(a, b noder.Hasher) bool { _ = "STUB: not implemented"; return false }

func (w *Worktree) Add(path string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) doAddDirectory(cfg *config.Config, idx *index.Index, s Status, directory string, ignorePattern []gitignore.Pattern) (added bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isPathInDirectory(path, directory string) bool { _ = "STUB: not implemented"; return false }

func (w *Worktree) AddWithOptions(opts *AddOptions) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) doAdd(path string, ignorePattern []gitignore.Pattern, skipStatus bool) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) AddGlob(pattern string) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) doAddFile(cfg *config.Config, idx *index.Index, s Status, path string, ignorePattern []gitignore.Pattern) (added bool, h plumbing.Hash, err error) {
	_ = "STUB: not implemented"
	return false, *new(plumbing.Hash), nil
}

func (w *Worktree) copyFileToStorage(cfg *config.Config, path string) (hash plumbing.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) fillEncodedObjectFromFile(cfg *config.Config, dst io.Writer, path string, _ os.FileInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) fillEncodedObjectFromSymlink(dst io.Writer, path string, _ os.FileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) addOrUpdateFileToIndex(idx *index.Index, filename string, h plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) doAddFileToIndex(idx *index.Index, filename string, h plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) doUpdateFileToIndex(e *index.Entry, filename string, h plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) Remove(path string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) doRemoveDirectory(idx *index.Index, directory string) (removed bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (w *Worktree) removeEmptyDirectory(path string) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) doRemoveFile(idx *index.Index, path string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) deleteFromIndex(idx *index.Index, path string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) deleteFromFilesystem(path string) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) RemoveGlob(pattern string) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) Move(from, to string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}
