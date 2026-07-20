package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

type LsRefsArgs struct {
	Peel        bool
	Symrefs     bool
	Unborn      bool
	RefPrefixes []string
}

func (r *LsRefsArgs) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func validateRefPrefix(p string) error { _ = "STUB: not implemented"; return nil }

const tooManyRefPrefixes = 65536

func (r *LsRefsArgs) Decode(rd io.Reader) error { _ = "STUB: not implemented"; return nil }

type LsRefsOutput struct {
	References []*plumbing.Reference
}

func (r *LsRefsOutput) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *LsRefsOutput) Decode(rd io.Reader) error { _ = "STUB: not implemented"; return nil }

func parseLsRefsLine(line string) ([]*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFullHash(s string) (plumbing.Hash, bool) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), false
}
