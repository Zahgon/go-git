package gitattributes

type Matcher interface {
	Match(path, attributes []string) (map[string]Attribute, bool)
}

type MatcherOptions struct{}

func NewMatcher(stack []MatchAttribute) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

type matcher struct {
	stack  []MatchAttribute
	macros map[string]MatchAttribute
}

func (m *matcher) init() {
	m.macros = make(map[string]MatchAttribute)

	for _, attr := range m.stack {
		if attr.Pattern == nil {
			m.macros[attr.Name] = attr
		}
	}
}

func (m *matcher) Match(path, attributes []string) (results map[string]Attribute, matched bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *matcher) expandMacro(name string, results map[string]Attribute) bool {
	_ = "STUB: not implemented"
	return false
}
