package gitignore

type MatchResult int

const (
	NoMatch MatchResult = iota

	Exclude

	Include
)

const (
	inclusionPrefix = "!"
	zeroToManyDirs  = "**"
	patternDirSep   = "/"
)

type Pattern interface {
	Match(path []string, isDir bool) MatchResult
}

type pattern struct {
	domain    []string
	pattern   []string
	inclusion bool
	dirOnly   bool
	isGlob    bool
}

func ParsePattern(p string, domain []string) Pattern {
	_ = "STUB: not implemented"
	return *new(Pattern)
}

func (p *pattern) Match(path []string, isDir bool) MatchResult {
	_ = "STUB: not implemented"
	return *new(MatchResult)
}

const (
	wmMatch           = 0
	wmNoMatch         = 1
	wmAbortAll        = -1
	wmAbortToStarStar = -2
)

const (
	wmCasefold = 1
	wmPathname = 2
)

func wildmatch(pattern, text string) bool { _ = "STUB: not implemented"; return false }

func dowild(p, text string, flags int) int { _ = "STUB: not implemented"; return 0 }

func isGlobSpecial(c byte) bool { _ = "STUB: not implemented"; return false }

func matchPOSIXClass(name string, ch byte, flags int) (matched, valid bool) {
	_ = "STUB: not implemented"
	return false, false
}

func isASCIIAlpha(ch byte) bool { _ = "STUB: not implemented"; return false }

func isASCIIDigit(ch byte) bool { _ = "STUB: not implemented"; return false }

func isASCIIUpper(ch byte) bool { _ = "STUB: not implemented"; return false }

func isASCIILower(ch byte) bool { _ = "STUB: not implemented"; return false }

func isASCIIPunct(ch byte) bool { _ = "STUB: not implemented"; return false }

func (p *pattern) simpleNameMatch(path []string, isDir bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *pattern) globMatch(path []string, isDir bool) bool {
	_ = "STUB: not implemented"
	return false
}
