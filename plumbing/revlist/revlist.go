package revlist

import (
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type objectWalker interface {
	RevListObjects(wants, haves []plumbing.Hash) ([]plumbing.Hash, error)
}

func Objects(
	s storer.EncodedObjectStorer,
	wants,
	haves []plumbing.Hash,
) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ObjectsWithRef(
	s storer.EncodedObjectStorer,
	wants,
	haves []plumbing.Hash,
) (map[plumbing.Hash][]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
