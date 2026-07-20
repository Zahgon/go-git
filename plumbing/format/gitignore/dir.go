package gitignore

import (
	"github.com/go-git/go-billy/v6"
)

const (
	commentPrefix   = "#"
	coreSection     = "core"
	excludesfile    = "excludesfile"
	gitDir          = ".git"
	gitignoreFile   = ".gitignore"
	gitconfigFile   = ".gitconfig"
	systemFile      = "/etc/gitconfig"
	infoExcludeFile = gitDir + "/info/exclude"
)

func readIgnoreFile(fs billy.Filesystem, path []string, ignoreFile string) (ps []Pattern, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadPatterns(fs billy.Filesystem, path []string) (ps []Pattern, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadPatterns(fs billy.Filesystem, path string) (ps []Pattern, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadGlobalPatterns(fs billy.Filesystem) (ps []Pattern, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadSystemPatterns(fs billy.Filesystem) (ps []Pattern, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
