package transactional

import "github.com/go-git/go-git/v6/config"

type ConfigStorage struct {
	config.ConfigStorer
	temporal config.ConfigStorer

	set bool
}

func NewConfigStorage(s, temporal config.ConfigStorer) *ConfigStorage {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigStorage) SetConfig(cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

func (c *ConfigStorage) Config() (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigStorage) Commit() error { _ = "STUB: not implemented"; return nil }
