package config

import (
	"io"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/config"
)

const (
	envGitConfigGlobal   = "GIT_CONFIG_GLOBAL"
	envGitConfigSystem   = "GIT_CONFIG_SYSTEM"
	envGitConfigNoSystem = "GIT_CONFIG_NOSYSTEM"
	envXDGConfigHome     = "XDG_CONFIG_HOME"
)

const maxConfigFileSize = 10 << 20

type Option func(*auto)

func WithFilesystem(fs billy.Basic) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewAuto(opts ...Option) *auto {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

type auto struct {
	fs billy.Basic
}

func (a *auto) Load(scope config.Scope) (config.ConfigStorer, error) {
	_ = "STUB: not implemented"
	return *new(config.ConfigStorer), nil
}

func (a *auto) loadGlobal() (*config.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *auto) loadSystem() (*config.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *auto) globalPaths() []string { _ = "STUB: not implemented"; return nil }

func xdgConfigPath(home string) string { _ = "STUB: not implemented"; return "" }

func systemPaths() []string { _ = "STUB: not implemented"; return nil }

func (a *auto) loadAndMerge(paths []string) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readAndClose(r io.ReadCloser) (cfg *config.Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isNoSystem() bool { _ = "STUB: not implemented"; return false }
