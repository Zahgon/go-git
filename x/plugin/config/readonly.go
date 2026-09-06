package config

import (
	"errors"

	"github.com/go-git/go-git/v6/config"
)

var ErrReadOnly = errors.New("config storer is read-only")

type readOnlyStorer struct {
	cfg config.Config
}

func (s *readOnlyStorer) Config() (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *readOnlyStorer) SetConfig(*config.Config) error { _ = "STUB: not implemented"; return nil }
