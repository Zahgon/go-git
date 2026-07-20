package ioutil

import (
	"context"
	"errors"
	"io"
)

type Peeker interface {
	Peek(int) ([]byte, error)
}

type ReadPeeker interface {
	io.Reader
	Peeker
}

var ErrEmptyReader = errors.New("reader is empty")

func NonEmptyReader(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

type readCloser struct {
	io.Reader
	closer io.Closer
}

func (r *readCloser) Close() error { _ = "STUB: not implemented"; return nil }

func NewReadCloser(r io.Reader, c io.Closer) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

type readCloserCloser struct {
	io.ReadCloser
	closer func() error
}

func (r *readCloserCloser) Close() (err error) { _ = "STUB: not implemented"; return nil }

func NewReadCloserWithCloser(r io.ReadCloser, c func() error) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

type writeCloser struct {
	io.Writer
	closer io.Closer
}

func (r *writeCloser) Close() error { _ = "STUB: not implemented"; return nil }

func NewWriteCloser(w io.Writer, c io.Closer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

type writeNopCloser struct {
	io.Writer
}

func (writeNopCloser) Close() error { _ = "STUB: not implemented"; return nil }

func WriteNopCloser(w io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

type readerAtAsReader struct {
	io.ReaderAt
	offset int64
}

func (r *readerAtAsReader) Read(bs []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func NewReaderUsingReaderAt(r io.ReaderAt, offset int64) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func CheckClose(c io.Closer, err *error) { _ = "STUB: not implemented"; return }

func NewContextWriteCloser(ctx context.Context, w io.WriteCloser) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func NewContextReadCloser(ctx context.Context, r io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

type readerOnError struct {
	io.Reader
	notify func(error)
}

func NewReaderOnError(r io.Reader, notify func(error)) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func NewReadCloserOnError(r io.ReadCloser, notify func(error)) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (r *readerOnError) Read(buf []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type writerOnError struct {
	io.Writer
	notify func(error)
}

func NewWriterOnError(w io.Writer, notify func(error)) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func NewWriteCloserOnError(w io.WriteCloser, notify func(error)) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (r *writerOnError) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type CloserFunc func() error

var _ io.Closer = CloserFunc(nil)

func (f CloserFunc) Close() error { _ = "STUB: not implemented"; return nil }
