package git

import (
	"syscall"
	"time"

	"github.com/go-git/go-git/v6/plumbing/format/index"
)

func init() {
	fillSystemInfo = func(e *index.Entry, sys any) {
		if os, ok := sys.(*syscall.Dir); ok {

			e.CreatedAt = time.Unix(int64(os.Mtime), 0)

			e.Dev = uint32(os.Dev)

			e.Inode = uint32(os.Qid.Path)

			e.GID = 0
			e.UID = 0
		}
	}
}

func isSymlinkWindowsNonAdmin(error) bool { _ = "STUB: not implemented"; return false }
