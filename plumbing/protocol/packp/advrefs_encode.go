package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

func (a *AdvRefs) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (a *AdvRefs) firstRef() (string, plumbing.Hash) {
	_ = "STUB: not implemented"
	return "", *new(plumbing.Hash)
}
