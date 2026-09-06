package dotgit

import (
	"github.com/go-git/go-git/v6/plumbing"
)

func (d *DotGit) setRef(fileName, content string, old *plumbing.Reference) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) setRefRwfs(fileName, content string, old *plumbing.Reference) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) setRefNorwfs(fileName, content string, old *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}
