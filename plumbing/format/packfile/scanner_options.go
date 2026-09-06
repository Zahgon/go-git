package packfile

import (
	"bufio"
)

type ScannerOption func(*Scanner)

func WithSHA256() ScannerOption { _ = "STUB: not implemented"; return *new(ScannerOption) }

func WithBufioReader(buf *bufio.Reader) ScannerOption {
	_ = "STUB: not implemented"
	return *new(ScannerOption)
}
