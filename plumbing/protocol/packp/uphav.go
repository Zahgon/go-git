package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

type UploadHaves struct {
	Haves []plumbing.Hash
	Done  bool
}

func (u *UploadHaves) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (u *UploadHaves) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }
