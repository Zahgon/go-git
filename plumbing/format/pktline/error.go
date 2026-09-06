package pktline

import (
	"errors"
	"io"
)

var (
	ErrInvalidErrorLine = errors.New("expected an error-line")

	ErrNilWriter = errors.New("nil writer")

	ErrNilReader = errors.New("nil reader")

	ErrNilError = errors.New("nil error")

	errPrefix = []byte("ERR ")
)

const (
	errPrefixSize = LenSize
)

type ErrorLine struct {
	Text string
}

func (e *ErrorLine) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorLine) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (e *ErrorLine) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }
