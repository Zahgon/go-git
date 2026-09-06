package filesystem

import (
	"github.com/go-git/go-git/v6/config"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

type ConfigStorage struct {
	dir          *dotgit.DotGit
	objectFormat formatcfg.ObjectFormat
}

func (c *ConfigStorage) Config() (conf *config.Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigStorage) SetConfig(cfg *config.Config) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigStorage) setWorktreeConfig(cfg *config.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigStorage) readBaseConfig() (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rawDiff(base, updated *formatcfg.Config) *formatcfg.Config {
	_ = "STUB: not implemented"
	return nil
}

func diffSubsections(
	baseSec *formatcfg.Section,
	updated formatcfg.Subsections,
) formatcfg.Subsections {
	_ = "STUB: not implemented"
	return *new(formatcfg.Subsections)
}

func diffOptions(
	baseOpts, updated formatcfg.Options,
) formatcfg.Options {
	_ = "STUB: not implemented"
	return *new(formatcfg.Options)
}

func baseOptions(sec *formatcfg.Section) formatcfg.Options {
	_ = "STUB: not implemented"
	return *new(formatcfg.Options)
}
