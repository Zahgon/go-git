package git

import (
	"context"
	"errors"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/index"
)

var (
	ErrSubmoduleAlreadyInitialized = errors.New("submodule already initialized")
	ErrSubmoduleNotInitialized     = errors.New("submodule not initialized")
)

type Submodule struct {
	initialized bool

	c *config.Submodule
	w *Worktree
}

func (s *Submodule) Config() *config.Submodule { _ = "STUB: not implemented"; return nil }

func (s *Submodule) Init() error { _ = "STUB: not implemented"; return nil }

func (s *Submodule) Status() (*SubmoduleStatus, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Submodule) status(idx *index.Index) (*SubmoduleStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Submodule) Repository() (*Repository, error) { _ = "STUB: not implemented"; return nil, nil }

func defaultRemote(r *Repository) (*config.RemoteConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lookupRemote(cfg *config.Config, name string) (*config.RemoteConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Submodule) Update(o *SubmoduleUpdateOptions) error { _ = "STUB: not implemented"; return nil }

func (s *Submodule) UpdateContext(ctx context.Context, o *SubmoduleUpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Submodule) update(ctx context.Context, o *SubmoduleUpdateOptions, forceHash plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Submodule) doRecursiveUpdate(ctx context.Context, r *Repository, o *SubmoduleUpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Submodule) fetchAndCheckout(
	ctx context.Context, r *Repository, o *SubmoduleUpdateOptions, hash plumbing.Hash,
) error {
	_ = "STUB: not implemented"
	return nil
}

type Submodules []*Submodule

func (s Submodules) Init() error { _ = "STUB: not implemented"; return nil }

func (s Submodules) Update(o *SubmoduleUpdateOptions) error { _ = "STUB: not implemented"; return nil }

func (s Submodules) UpdateContext(ctx context.Context, o *SubmoduleUpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (s Submodules) Status() (SubmodulesStatus, error) {
	_ = "STUB: not implemented"
	return *new(SubmodulesStatus), nil
}

type SubmodulesStatus []*SubmoduleStatus

func (s SubmodulesStatus) String() string { _ = "STUB: not implemented"; return "" }

type SubmoduleStatus struct {
	Path     string
	Current  plumbing.Hash
	Expected plumbing.Hash
	Branch   plumbing.ReferenceName
}

func (s *SubmoduleStatus) IsClean() bool { _ = "STUB: not implemented"; return false }

func (s *SubmoduleStatus) String() string { _ = "STUB: not implemented"; return "" }
