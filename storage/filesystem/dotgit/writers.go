package dotgit

import (
	"hash"
	"io"
	"sync/atomic"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
	"github.com/go-git/go-git/v6/plumbing/format/objfile"
	"github.com/go-git/go-git/v6/plumbing/format/packfile"
)

type PackWriter struct {
	Notify func(plumbing.Hash, *idxfile.Writer)

	fs       billy.Filesystem
	fr, fw   billy.File
	synced   *syncedReader
	checksum plumbing.Hash
	parser   *packfile.Parser
	writer   *idxfile.Writer
	result   chan error
	format   formatcfg.ObjectFormat
	writeRev bool
}

func newPackWrite(fs billy.Filesystem, format formatcfg.ObjectFormat, writeRev bool) (*PackWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *PackWriter) buildIndex() { _ = "STUB: not implemented"; return }

func (w *PackWriter) waitBuildIndex() error { _ = "STUB: not implemented"; return nil }

func (w *PackWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *PackWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w *PackWriter) clean() error { _ = "STUB: not implemented"; return nil }

func (w *PackWriter) save() error { _ = "STUB: not implemented"; return nil }

func fileExists(fs billy.Filesystem, path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (w *PackWriter) encodeIdx(writer io.Writer, h hash.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *PackWriter) encodeRev(writer io.Writer, h hash.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

type syncedReader struct {
	w io.Writer
	r io.ReadSeeker

	blocked, done atomic.Uint32
	written, read atomic.Uint64
	news          chan bool
}

func newSyncedReader(w io.Writer, r io.ReadSeeker) *syncedReader {
	_ = "STUB: not implemented"
	return nil
}

func (s *syncedReader) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *syncedReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *syncedReader) isDone() bool { _ = "STUB: not implemented"; return false }

func (s *syncedReader) isBlocked() bool { _ = "STUB: not implemented"; return false }

func (s *syncedReader) wake() { _ = "STUB: not implemented"; return }

func (s *syncedReader) sleep() { _ = "STUB: not implemented"; return }

func (s *syncedReader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *syncedReader) Close() error { _ = "STUB: not implemented"; return nil }

type ObjectWriter struct {
	objfile.Writer
	fs billy.Filesystem
	f  billy.File
}

func newObjectWriter(fs billy.Filesystem, objectFormat formatcfg.ObjectFormat) (*ObjectWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *ObjectWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w *ObjectWriter) save() error { _ = "STUB: not implemented"; return nil }
