package git

import (
	"errors"
	"regexp"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/storage"
)

var (
	ErrEmptyCommit = errors.New("cannot create empty commit: clean working tree")

	ErrCannotCherryPickWithoutCommitOptions = errors.New("cannot cherry-pick without commit options")

	invalidCharactersRe = regexp.MustCompile(`[<>\n]`)
)

func (w *Worktree) Commit(msg string, opts *CommitOptions) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) CherryPick(commitOpts *CommitOptions, ortStrategyOption OrtMergeStrategyOption, commits ...*object.Commit) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worktree) autoAddModifiedAndDeleted() error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) updateHEAD(commit plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (w *Worktree) buildCommitObject(msg string, opts *CommitOptions, tree plumbing.Hash) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (w *Worktree) sanitize(signature object.Signature) object.Signature {
	_ = "STUB: not implemented"
	return *new(object.Signature)
}

type buildTreeHelper struct {
	fs billy.Filesystem
	s  storage.Storer

	trees   map[string]*object.Tree
	entries map[string]*object.TreeEntry
}

func (h *buildTreeHelper) BuildTree(idx *index.Index, _ *CommitOptions) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (h *buildTreeHelper) commitIndexEntry(e *index.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *buildTreeHelper) doBuildTree(e *index.Entry, parent, fullpath string) {
	_ = "STUB: not implemented"
	return
}

type sortableEntries []object.TreeEntry

func (sortableEntries) sortName(te object.TreeEntry) string { _ = "STUB: not implemented"; return "" }

func (se sortableEntries) Len() int           { _ = "STUB: not implemented"; return 0 }
func (se sortableEntries) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (se sortableEntries) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *buildTreeHelper) copyTreeToStorageRecursive(parent string, t *object.Tree) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}
