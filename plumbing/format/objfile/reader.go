package objfile

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	format "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/utils/sync"
)

var (
	ErrClosed        = errors.New("objfile: already closed")
	ErrHeader        = errors.New("objfile: invalid header")
	ErrHeaderTooLong = errors.New("objfile: header exceeds maximum length")
	ErrHeaderNotRead = errors.New("objfile: Header must be called before Read")
	ErrNegativeSize  = errors.New("objfile: negative object size")
)

const maxHeaderLen = 32

type Reader struct {
	multi        io.Reader
	zlib         *sync.ZLibReader
	hasher       plumbing.Hasher
	objectFormat format.ObjectFormat
	closed       bool
}

func NewReader(r io.Reader, objectFormat format.ObjectFormat) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) Header() (t plumbing.ObjectType, size int64, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType), 0, nil
}

func (r *Reader) readUntil(delim byte, budget int) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (r *Reader) prepareForRead(t plumbing.ObjectType, size int64) {
	_ = "STUB: not implemented"
	return
}

func (r *Reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) Hash() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }
