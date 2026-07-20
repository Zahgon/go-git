package packp

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

var (
	ErrEmptyAdvRefs = errors.New("empty advertised-ref message")

	ErrEmptyInput = errors.New("empty input")
)

func (a *AdvRefs) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func hashFrom(line []byte) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func parseRef(data []byte) (string, plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return "", *new(plumbing.Hash), nil
}
