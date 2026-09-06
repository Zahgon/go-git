package gitattributes

const (
	patternDirSep  = "/"
	zeroToManyDirs = "**"
)

type Pattern interface {
	Match(path []string) bool
}

type pattern struct {
	domain  []string
	pattern []string
}

func ParsePattern(p string, domain []string) Pattern {
	_ = "STUB: not implemented"
	return *new(Pattern)
}

func (p *pattern) Match(path []string) bool { _ = "STUB: not implemented"; return false }
