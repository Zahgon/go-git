package storage

import (
	"errors"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

var ErrReferenceHasChanged = errors.New("reference has changed concurrently")

type Storer interface {
	storer.EncodedObjectStorer
	storer.ReferenceStorer
	storer.ShallowStorer
	storer.IndexStorer
	config.ConfigStorer
	ModuleStorer
}

type ModuleStorer interface {
	Module(name string) (Storer, error)
}
