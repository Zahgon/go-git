package packfile

import (
	"bufio"
	"io"
)

type scannerReader struct {
	reader io.Reader
	crc    io.Writer
	rbuf   *bufio.Reader
	wbuf   *bufio.Writer
	offset int64
	seeker io.Seeker
}

func newScannerReader(r io.Reader, h io.Writer, rbuf *bufio.Reader) *scannerReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *scannerReader) Reset(reader io.Reader) { _ = "STUB: not implemented"; return }

func (r *scannerReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *scannerReader) ReadByte() (b byte, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *scannerReader) Flush() error { _ = "STUB: not implemented"; return nil }

func (r *scannerReader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
