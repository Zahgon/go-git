package packp

import (
	"errors"
	"io"
)

var ErrInvalidPushOption = errors.New("invalid push option")

type PushOptions struct {
	Options []string
}

func (opts *PushOptions) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (opts *PushOptions) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func isNotGraphic(r rune) bool { _ = "STUB: not implemented"; return false }
