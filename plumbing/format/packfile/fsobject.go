package packfile

import (
	"bufio"
	"io"
	stdsync "sync"

	billy "github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
	"github.com/go-git/go-git/v6/utils/sync"
)

const probeSize = 1

var probeBufPool = stdsync.Pool{
	New: func() any {
		var buf [probeSize]byte
		return &buf
	},
}

func probePack(pack io.ReaderAt, offset int64) error { _ = "STUB: not implemented"; return nil }

type FSObject struct {
	hash     plumbing.Hash
	offset   int64
	size     int64
	typ      plumbing.ObjectType
	index    idxfile.Index
	fs       billy.Filesystem
	pack     billy.File
	packPath string
	cache    cache.Object

	acquireRandom func() (RandomReader, error)
}

func NewFSObject(
	hash plumbing.Hash,
	finalType plumbing.ObjectType,
	offset int64,
	contentSize int64,
	index idxfile.Index,
	fs billy.Filesystem,
	pack billy.File,
	packPath string,
	cache cache.Object,
) *FSObject {
	_ = "STUB: not implemented"
	return nil
}

func (o *FSObject) Reader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type zlibReadCloser struct {
	r      *sync.ZLibReader
	f      io.Closer
	rbuf   *bufio.Reader
	closed bool
}

func (r *zlibReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *zlibReadCloser) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (o *FSObject) SetSize(int64) { _ = "STUB: not implemented"; return }

func (o *FSObject) SetType(plumbing.ObjectType) { _ = "STUB: not implemented"; return }

func (o *FSObject) Hash() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

func (o *FSObject) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (o *FSObject) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (o *FSObject) Writer() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}
