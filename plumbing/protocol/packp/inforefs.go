package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

type InfoRefs struct {
	References []*plumbing.Reference
}

func (i *InfoRefs) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (i *InfoRefs) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }
