package packfile

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

type PackHandle interface {
	OpenPackReader() (io.ReadSeekCloser, error)

	OpenRandomReader() (RandomReader, error)

	PackHash() (plumbing.Hash, error)
}

type RandomReader interface {
	io.ReaderAt
	io.Closer
}

type PackHandleResolver func() (PackHandle, error)

func WithPackHandle(get PackHandleResolver) PackfileOption {
	_ = "STUB: not implemented"
	return *new(PackfileOption)
}

func (p *Packfile) openRandomReader() (RandomReader, error) {
	_ = "STUB: not implemented"
	return *new(RandomReader), nil
}
