package packp

import (
	"errors"
	"io"
)

var ErrInvalidSmartReply = errors.New("invalid smart reply")

type SmartReply struct {
	Service string
}

func (s *SmartReply) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (s *SmartReply) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }
