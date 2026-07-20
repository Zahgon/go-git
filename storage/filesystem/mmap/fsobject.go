//go:build darwin || linux

package mmap

import (
	"bufio"
	"io"
	gosync "sync"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/utils/sync"
)

const (
	maskContinue = 0x80
)

type ondemandObject struct {
	hash        plumbing.Hash
	offset      int64
	size        int64
	typ         plumbing.ObjectType
	scanner     *PackScanner
	diskType    plumbing.ObjectType
	autoResolve bool

	m gosync.RWMutex
}

func newOndemandObject(
	hash plumbing.Hash,
	typ plumbing.ObjectType,
	offset int64,
	size int64,
	scanner *PackScanner,
	autoResolve bool,
) *ondemandObject {
	_ = "STUB: not implemented"
	return nil
}

func (o *ondemandObject) Resolve() error { _ = "STUB: not implemented"; return nil }

func (o *ondemandObject) Reader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (o *ondemandObject) Hash() plumbing.Hash {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash)
}

func (o *ondemandObject) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (o *ondemandObject) Type() plumbing.ObjectType {
	_ = "STUB: not implemented"
	return *new(plumbing.ObjectType)
}

func (o *ondemandObject) SetSize(int64) { _ = "STUB: not implemented"; return }

func (o *ondemandObject) SetType(plumbing.ObjectType) { _ = "STUB: not implemented"; return }

func (o *ondemandObject) Writer() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (o *ondemandObject) resolveMetadata() error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck
//nolint:ineffassign

func (o *ondemandObject) toDataOffset(offset int64) int64 { _ = "STUB: not implemented"; return 0 }

func (o *ondemandObject) resolveDelta() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type zlibReadCloser struct {
	r        *sync.ZLibReader
	rbuf     *bufio.Reader
	once     gosync.Once
	closeErr error
}

func (r *zlibReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *zlibReadCloser) Close() (err error) { _ = "STUB: not implemented"; return nil }
