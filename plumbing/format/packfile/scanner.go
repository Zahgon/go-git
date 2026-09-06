package packfile

import (
	"bufio"
	"errors"
	"hash"
	"io"
	"sync"

	"github.com/go-git/go-git/v6/plumbing"
	gogithash "github.com/go-git/go-git/v6/plumbing/hash"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

var (
	ErrEmptyPackfile = NewError("empty packfile")

	ErrBadSignature = NewError("bad signature")

	ErrMalformedPackfile = NewError("malformed pack file")

	ErrUnsupportedVersion = NewError("unsupported packfile version")

	ErrSeekNotSupported = NewError("not seek support")

	ErrInflatedSizeMismatch = errors.New("packfile: inflated object exceeds declared size")
)

type boundedWriter struct {
	w     io.Writer
	limit int64
	n     int64
}

func (b *boundedWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type BoundedReadCloser struct {
	lr      io.LimitedReader
	closer  io.Closer
	overrun bool
}

func NewBoundedReadCloser(rc io.ReadCloser, limit int64) *BoundedReadCloser {
	_ = "STUB: not implemented"
	return nil
}

func (b *BoundedReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *BoundedReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

type Scanner struct {
	version Version

	objects uint32

	objIndex int

	hasher plumbing.Hasher

	crc hash.Hash32

	packhash gogithash.Hash

	objectIDSize int

	nextFn stateFn

	packData PackData

	err error

	m sync.Mutex

	storage storer.EncodedObjectStorer

	*scannerReader
	rbuf *bufio.Reader

	lowMemoryMode bool
}

func NewScanner(rs io.Reader, opts ...ScannerOption) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

func (r *Scanner) Scan() bool { _ = "STUB: not implemented"; return false }

func (r *Scanner) Reset() error { _ = "STUB: not implemented"; return nil }

func (r *Scanner) Data() PackData { _ = "STUB: not implemented"; return *new(PackData) }

func (r *Scanner) Error() error { _ = "STUB: not implemented"; return nil }

func (r *Scanner) SeekFromStart(offset int64) error { _ = "STUB: not implemented"; return nil }

func (r *Scanner) WriteObject(oh *ObjectHeader, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Scanner) inflateContent(contentOffset int64, writer io.Writer, declaredSize int64) error {
	_ = "STUB: not implemented"
	return nil
}

func scan(r *Scanner) error { _ = "STUB: not implemented"; return nil }

type stateFn func(*Scanner) (stateFn, error)

func packHeaderSignature(r *Scanner) (stateFn, error) {
	_ = "STUB: not implemented"
	return *new(stateFn), nil
}

func packVersion(r *Scanner) (stateFn, error) { _ = "STUB: not implemented"; return *new(stateFn), nil }

func packObjectsQty(r *Scanner) (stateFn, error) {
	_ = "STUB: not implemented"
	return *new(stateFn), nil
}

func objectEntry(r *Scanner) (stateFn, error) { _ = "STUB: not implemented"; return *new(stateFn), nil }

func packFooter(r *Scanner) (stateFn, error) { _ = "STUB: not implemented"; return *new(stateFn), nil }
