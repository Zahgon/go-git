//go:build windows

package dotgit

import (
	"github.com/go-git/go-billy/v6"
)

func fixPermissions(fs billy.Filesystem, path string) { _ = "STUB: not implemented"; return }

func isReadOnly(fs billy.Filesystem, path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
