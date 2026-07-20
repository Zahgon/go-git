package config

import (
	"errors"
	"regexp"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

var (
	ErrModuleEmptyURL = errors.New("module config: empty URL")

	ErrModuleEmptyPath = errors.New("module config: empty path")

	ErrModuleBadPath = errors.New("submodule has an invalid path")

	ErrModuleBadName = errors.New("ignoring suspicious submodule name")
)

var dotdotPath = regexp.MustCompile(`(^|[/\\])\.\.([/\\]|$)`)

type Modules struct {
	Submodules map[string]*Submodule

	raw *format.Config
}

func NewModules() *Modules { _ = "STUB: not implemented"; return nil }

const (
	pathKey   = "path"
	branchKey = "branch"
)

func (m *Modules) Unmarshal(b []byte) error { _ = "STUB: not implemented"; return nil }

func (m *Modules) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type Submodule struct {
	Name string

	Path string

	URL string

	Branch string

	raw *format.Subsection
}

func (m *Submodule) Validate() error { _ = "STUB: not implemented"; return nil }

func validSubmoduleName(name string) error { _ = "STUB: not implemented"; return nil }

func isPathSep(r rune) bool { _ = "STUB: not implemented"; return false }

func (m *Submodule) unmarshal(s *format.Subsection) { _ = "STUB: not implemented"; return }

func (m *Submodule) marshal() *format.Subsection { _ = "STUB: not implemented"; return nil }
