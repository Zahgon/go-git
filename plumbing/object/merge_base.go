package object

import (
	"fmt"

	"github.com/go-git/go-git/v6/plumbing"
)

var errIsReachable = fmt.Errorf("first is reachable from second")

func (c *Commit) MergeBase(other *Commit) ([]*Commit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Commit) IsAncestor(other *Commit) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ancestorsIndex(excluded, starting *Commit) (map[plumbing.Hash]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Independents(commits []*Commit) ([]*Commit, error) { _ = "STUB: not implemented"; return nil, nil }

func sortByCommitDateDesc(commits ...*Commit) []*Commit { _ = "STUB: not implemented"; return nil }

func indexOf(commits []*Commit, target *Commit) int { _ = "STUB: not implemented"; return 0 }

func remove(commits []*Commit, toDelete *Commit) []*Commit { _ = "STUB: not implemented"; return nil }

func removeDuplicated(commits []*Commit) []*Commit { _ = "STUB: not implemented"; return nil }

func isInIndexCommitFilter(index map[plumbing.Hash]struct{}) CommitFilter {
	_ = "STUB: not implemented"
	return *new(CommitFilter)
}
