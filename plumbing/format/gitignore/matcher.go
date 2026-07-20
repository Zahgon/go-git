package gitignore

type Matcher interface {
	Match(path []string, isDir bool) bool
}

func NewMatcher(ps []Pattern) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

type matcher struct {
	patterns []Pattern
}

func (m *matcher) Match(path []string, isDir bool) bool { _ = "STUB: not implemented"; return false }
