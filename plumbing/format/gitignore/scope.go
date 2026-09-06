package gitignore

import (
	"github.com/go-git/go-billy/v6"
)

type Scope struct {
	patterns []Pattern
	matcher  Matcher

	excluded bool
}

func NewScope(base []Pattern) *Scope { _ = "STUB: not implemented"; return nil }

func (s *Scope) Descend(dir []string, readOwn func() ([]Pattern, error)) (*Scope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scope) Excluded() bool { _ = "STUB: not implemented"; return false }

func (s *Scope) Match(path []string, isDir bool) bool { _ = "STUB: not implemented"; return false }

func (s *Scope) Patterns() []Pattern { _ = "STUB: not implemented"; return nil }

func (s *Scope) matches(path []string, isDir bool) bool { _ = "STUB: not implemented"; return false }

func DirPatterns(fs billy.Filesystem, path []string) ([]Pattern, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RootPatterns(fs billy.Filesystem) ([]Pattern, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
