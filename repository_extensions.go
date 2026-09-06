package git

import (
	"errors"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/storage"
)

var (
	ErrUnsupportedExtensionRepositoryFormatVersion = errors.New("core.repositoryformatversion does not support extension")

	ErrUnsupportedRepositoryFormatVersion = errors.New("core.repositoryformatversion not supported")

	ErrUnknownExtension = errors.New("unknown extension")

	builtinExtensions = map[string]struct{}{

		"noop": {},

		"noop-v1": {},
	}

	extensionsValidForV0 = map[string]struct{}{
		"noop":            {},
		"partialclone":    {},
		"preciousobjects": {},
		"worktreeconfig":  {},
	}
)

type extension struct {
	name  string
	value string
}

func extensions(cfg *config.Config) []extension { _ = "STUB: not implemented"; return nil }

func verifyExtensions(st storage.Storer, cfg *config.Config) error {
	_ = "STUB: not implemented"
	return nil
}
