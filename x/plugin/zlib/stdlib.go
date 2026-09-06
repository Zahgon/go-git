package zlib

import (
	"io"
)

type Reader interface {
	io.ReadCloser
	Reset(r io.Reader, dict []byte) error
}

type Writer interface {
	io.WriteCloser
	Reset(w io.Writer)
	Flush() error
}

type Provider interface {
	NewReader(r io.Reader) (Reader, error)
	NewWriter(w io.Writer) Writer
}

type Stdlib struct{}

func NewStdlib() *Stdlib { _ = "STUB: not implemented"; return nil }

func (*Stdlib) NewReader(r io.Reader) (Reader, error) {
	_ = "STUB: not implemented"
	return *new(Reader), nil
}

func (*Stdlib) NewWriter(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }
