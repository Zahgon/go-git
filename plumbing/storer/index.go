package storer

import "github.com/go-git/go-git/v6/plumbing/format/index"

type IndexStorer interface {
	SetIndex(*index.Index) error
	Index() (*index.Index, error)
}
