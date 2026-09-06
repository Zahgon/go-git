package git

import (
	"context"
	"io"
	"net"

	"github.com/go-git/go-git/v6/plumbing/transport"
)

const DefaultPort = 9418

type Options struct {
	DialContext transport.DialContextFunc

	DialProxy func(transport.DialContextFunc) transport.DialContextFunc
}

type Transport struct {
	opts Options
}

func NewTransport(opts Options) *Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) Connect(ctx context.Context, req *transport.Request) (transport.Conn, error) {
	_ = "STUB: not implemented"
	return *new(transport.Conn), nil
}

type gitConn struct {
	nc net.Conn
}

var _ transport.Conn = (*gitConn)(nil)

func (c *gitConn) Reader() io.Reader      { _ = "STUB: not implemented"; return *new(io.Reader) }
func (c *gitConn) Writer() io.WriteCloser { _ = "STUB: not implemented"; return *new(io.WriteCloser) }
func (c *gitConn) Close() error           { _ = "STUB: not implemented"; return nil }
