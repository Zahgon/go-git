package diff

import (
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func Do(src, dst string) (diffs []diffmatchpatch.Diff) { _ = "STUB: not implemented"; return nil }

func DoWithTimeout(src, dst string, timeout time.Duration) (diffs []diffmatchpatch.Diff) {
	_ = "STUB: not implemented"
	return nil
}

func Dst(diffs []diffmatchpatch.Diff) string { _ = "STUB: not implemented"; return "" }

func Src(diffs []diffmatchpatch.Diff) string { _ = "STUB: not implemented"; return "" }
