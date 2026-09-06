package gitattributes

import (
	"github.com/go-git/go-billy/v6"
)

const (
	coreSection       = "core"
	attributesfile    = "attributesfile"
	gitDir            = ".git"
	gitattributesFile = ".gitattributes"
	gitconfigFile     = ".gitconfig"
	systemFile        = "/etc/gitconfig"
)

func ReadAttributesFile(fs billy.Filesystem, path []string, attributesFile string, allowMacro bool) ([]MatchAttribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadPatterns(fs billy.Filesystem, path []string) (attributes []MatchAttribute, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func walkDirectory(fs billy.Filesystem, root []string) (attributes []MatchAttribute, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadPatterns(fs billy.Filesystem, path string) ([]MatchAttribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadGlobalPatterns(fs billy.Filesystem) (attributes []MatchAttribute, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadSystemPatterns(fs billy.Filesystem) (attributes []MatchAttribute, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
