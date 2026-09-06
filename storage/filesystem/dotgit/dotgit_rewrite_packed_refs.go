package dotgit

import (
	"github.com/go-git/go-billy/v6"
)

func (d *DotGit) openAndLockPackedRefsMode() int { _ = "STUB: not implemented"; return 0 }

func (d *DotGit) rewritePackedRefsWhileLocked(
	tmp, pr billy.File,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) copyToExistingFile(tmp, pr billy.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DotGit) copyNewFile(tmp, pr billy.File) (err error) { _ = "STUB: not implemented"; return nil }
