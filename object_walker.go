package git

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/storage"
)

type objectWalker struct {
	Storer storage.Storer

	seen map[plumbing.Hash]struct{}
}

func newObjectWalker(s storage.Storer) *objectWalker { _ = "STUB: not implemented"; return nil }

func (p *objectWalker) walkAllRefs() error { _ = "STUB: not implemented"; return nil }

func (p *objectWalker) isSeen(hash plumbing.Hash) bool { _ = "STUB: not implemented"; return false }

func (p *objectWalker) add(hash plumbing.Hash) { _ = "STUB: not implemented"; return }

func (p *objectWalker) walkObjectTree(hash plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}
