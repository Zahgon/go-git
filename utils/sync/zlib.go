package sync

import (
	"errors"
	"io"
	"sync"

	"github.com/go-git/go-git/v6/x/plugin"
)

type ZlibReader = plugin.ZlibReader

type ZlibWriter = plugin.ZlibWriter

var errNilZlibReader = errors.New("utils/sync: zlib reader source is nil")

var (
	zlibInitBytes = []byte{0x78, 0x9c, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0x00, 0x01}

	zlibProviderOnce sync.Once
	zlibProvider     plugin.ZlibProvider

	zlibReader = sync.Pool{New: newPooledZlibReader}
	zlibWriter = sync.Pool{New: newPooledZlibWriter}
)

func getZlibProvider() plugin.ZlibProvider {
	_ = "STUB: not implemented"
	return *new(plugin.ZlibProvider)
}

func newPooledZlibReader() any { _ = "STUB: not implemented"; return *new(any) }

func newPooledZlibWriter() any { _ = "STUB: not implemented"; return *new(any) }

type ZLibReader struct {
	dict   *[]byte
	reader ZlibReader
}

func (r *ZLibReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *ZLibReader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *ZLibReader) Reset(in io.Reader, dict []byte) error { _ = "STUB: not implemented"; return nil }

func GetZlibReader(r io.Reader) (*ZLibReader, error) { _ = "STUB: not implemented"; return nil, nil }

func PutZlibReader(z *ZLibReader) { _ = "STUB: not implemented"; return }

func GetZlibWriter(w io.Writer) ZlibWriter { _ = "STUB: not implemented"; return *new(ZlibWriter) }

func PutZlibWriter(w ZlibWriter) { _ = "STUB: not implemented"; return }
