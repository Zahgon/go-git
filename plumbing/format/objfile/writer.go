package objfile

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	format "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/utils/sync"
)

var ErrOverflow = errors.New("objfile: declared data length exceeded (overflow)")

type Writer struct {
	raw          io.Writer
	hasher       plumbing.Hasher
	multi        io.Writer
	zlib         sync.ZlibWriter
	objectFormat format.ObjectFormat

	closed  bool
	pending int64

	closeErr error
}

func NewWriter(w io.Writer, objectFormat format.ObjectFormat) *Writer {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) WriteHeader(t plumbing.ObjectType, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) writeHeader(t plumbing.ObjectType, typeBytes []byte, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) prepareForWrite(t plumbing.ObjectType, size int64) {
	_ = "STUB: not implemented"
	return
}

func (w *Writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) Hash() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }
