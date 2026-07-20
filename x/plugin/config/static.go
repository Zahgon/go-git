package config

import (
	"github.com/go-git/go-git/v6/config"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
)

func NewStatic(global, system config.Config) *static {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

type static struct {
	global config.Config
	system config.Config
}

func (s *static) Load(scope config.Scope) (config.ConfigStorer, error) {
	_ = "STUB: not implemented"
	return *new(config.ConfigStorer), nil
}

func cloneConfig(c *config.Config) *config.Config { _ = "STUB: not implemented"; return nil }

func cloneRemotes(m map[string]*config.RemoteConfig) map[string]*config.RemoteConfig {
	_ = "STUB: not implemented"
	return nil
}

func cloneURLs(s []*config.URL) []*config.URL { _ = "STUB: not implemented"; return nil }

func cloneMapShallow[K comparable, V any](m map[K]*V) map[K]*V {
	_ = "STUB: not implemented"
	return nil
}

func cloneSlice[T any](s []T) []T { _ = "STUB: not implemented"; return nil }

func cloneRawConfig(c *formatcfg.Config) *formatcfg.Config { _ = "STUB: not implemented"; return nil }

func cloneRawSubsections(ss formatcfg.Subsections) formatcfg.Subsections {
	_ = "STUB: not implemented"
	return *new(formatcfg.Subsections)
}

func cloneRawOptions(opts formatcfg.Options) formatcfg.Options {
	_ = "STUB: not implemented"
	return *new(formatcfg.Options)
}
