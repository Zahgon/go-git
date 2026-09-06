package sideband

import (
	"io"
)

type Muxer struct {
	max int
	w   io.Writer
}

const chLen = 1

func NewMuxer(t Type, w io.Writer) *Muxer { _ = "STUB: not implemented"; return nil }

func (m *Muxer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Muxer) WriteChannel(t Channel, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Muxer) doWrite(ch Channel, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
