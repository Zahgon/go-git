package git

import (
	"errors"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
)

type (
	PruneHandler func(unreferencedObjectHash plumbing.Hash) error

	PruneOptions struct {
		OnlyObjectsOlderThan time.Time

		Handler PruneHandler
	}
)

var ErrLooseObjectsNotSupported = errors.New("loose objects not supported")

func (r *Repository) DeleteObject(hash plumbing.Hash) error { _ = "STUB: not implemented"; return nil }

func (r *Repository) Prune(opt PruneOptions) error { _ = "STUB: not implemented"; return nil }
