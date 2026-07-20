package file

import (
	"context"
	"io"

	"github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/storage"
)

type ServerFunc func(ctx context.Context, st storage.Storer, r io.ReadCloser, w io.WriteCloser, gitProtocol string) error

func defaultUploadPack(ctx context.Context, st storage.Storer, r io.ReadCloser, w io.WriteCloser, gitProtocol string) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultReceivePack(ctx context.Context, st storage.Storer, r io.ReadCloser, w io.WriteCloser, gitProtocol string) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultUploadArchive(ctx context.Context, st storage.Storer, r io.ReadCloser, w io.WriteCloser, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

type Options struct {
	Loader transport.Loader
}

type Transport struct {
	loader        transport.Loader
	uploadPack    ServerFunc
	receivePack   ServerFunc
	uploadArchive ServerFunc
}

func NewTransport(opts Options) *Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) Connect(ctx context.Context, req *transport.Request) (transport.Conn, error) {
	_ = "STUB: not implemented"
	return *new(transport.Conn), nil
}

func (t *Transport) connect(ctx context.Context, req *transport.Request) (io.Reader, *io.PipeWriter, func() error, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil, nil, nil
}

type fileConn struct {
	r     io.Reader
	w     io.WriteCloser
	close func() error
}

var _ transport.Conn = (*fileConn)(nil)

func (c *fileConn) Reader() io.Reader      { _ = "STUB: not implemented"; return *new(io.Reader) }
func (c *fileConn) Writer() io.WriteCloser { _ = "STUB: not implemented"; return *new(io.WriteCloser) }
func (c *fileConn) Close() error           { _ = "STUB: not implemented"; return nil }
