package convert

import (
	"io"
)

type crlfToLFWriter struct {
	w io.Writer
}

func NewLFWriter(w io.Writer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (conv *crlfToLFWriter) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type lfToCRLFWriter struct {
	w     io.Writer
	hadCR bool
}

func NewCRLFWriter(w io.Writer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (conv *lfToCRLFWriter) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
