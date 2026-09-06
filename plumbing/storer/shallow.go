package storer

import "github.com/go-git/go-git/v6/plumbing"

type ShallowStorer interface {
	SetShallow([]plumbing.Hash) error
	Shallow() ([]plumbing.Hash, error)
}
