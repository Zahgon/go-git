package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

const (
	shallowLineLen   = 48
	unshallowLineLen = 50
)

type ShallowUpdate struct {
	Shallows   []plumbing.Hash
	Unshallows []plumbing.Hash
}

func (r *ShallowUpdate) Decode(reader io.Reader) error { _ = "STUB: not implemented"; return nil }

func (r *ShallowUpdate) decodeShallowLine(line []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ShallowUpdate) decodeUnshallowLine(line []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShallowUpdate) decodeLine(line, prefix []byte, expLen int) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (r *ShallowUpdate) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }
