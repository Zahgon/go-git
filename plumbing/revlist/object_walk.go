package revlist

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type objectWalk struct {
	s          storer.EncodedObjectStorer
	shallows   map[plumbing.Hash]struct{}
	wantsQueue []*object.Commit
	havesQueue []*object.Commit
	wantsSeen  map[plumbing.Hash]struct{}
	havesSeen  map[plumbing.Hash]struct{}
	seen       map[plumbing.Hash]struct{}
	result     []plumbing.Hash
}

func newObjectWalk(s storer.EncodedObjectStorer) (*objectWalk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func shallowSet(s storer.EncodedObjectStorer) (map[plumbing.Hash]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *objectWalk) seedWants(wants []plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (w *objectWalk) seedHaves(haves []plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

const (
	wantPaint uint8 = 1 << iota
	havePaint
)

func (w *objectWalk) walk() error { _ = "STUB: not implemented"; return nil }

type missingParent struct {
	hash  plumbing.Hash
	child plumbing.Hash
}

func (w *objectWalk) propagate(queue *[]*object.Commit, flags map[plumbing.Hash]uint8, missing *[]missingParent, lc *object.Commit, f uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func allStale(queue []*object.Commit, flags map[plumbing.Hash]uint8) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *objectWalk) walkFull() error { _ = "STUB: not implemented"; return nil }

func (w *objectWalk) processCommitTrees(lc *object.Commit) error {
	_ = "STUB: not implemented"
	return nil
}

func insertSorted(q *[]*object.Commit, c *object.Commit) { _ = "STUB: not implemented"; return }

func collectChangedTreeObjects(
	s storer.EncodedObjectStorer,
	newTree *object.Tree,
	oldTrees []*object.Tree,
	seen map[plumbing.Hash]struct{},
	result *[]plumbing.Hash,
) error {
	_ = "STUB: not implemented"
	return nil
}

func collectAllTreeObjects(
	s storer.EncodedObjectStorer,
	t *object.Tree,
	seen map[plumbing.Hash]struct{},
	result *[]plumbing.Hash,
) error {
	_ = "STUB: not implemented"
	return nil
}

func markTreeSeen(s storer.EncodedObjectStorer, t *object.Tree, seen map[plumbing.Hash]struct{}) {
	_ = "STUB: not implemented"
	return
}
