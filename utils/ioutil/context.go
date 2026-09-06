package ioutil

import (
	"context"
	"io"
)

type ioret struct {
	err error
	n   int
}

type Writer interface {
	io.Writer
}

type ctxWriter struct {
	w   io.Writer
	ctx context.Context
}

func NewContextWriter(ctx context.Context, w io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (w *ctxWriter) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type Reader interface {
	io.Reader
}

type ctxReader struct {
	r      io.Reader
	ctx    context.Context
	closer io.Closer
}

func NewContextReader(ctx context.Context, r io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (r *ctxReader) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func NewContextReaderWithCloser(ctx context.Context, r io.Reader, closer io.Closer) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}
