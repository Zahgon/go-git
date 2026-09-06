package transport

import (
	"context"
	"io"
)

type Conn interface {
	io.Closer
	Reader() io.Reader
	Writer() io.WriteCloser
}

type Connector interface {
	Connect(context.Context, *Request) (Conn, error)
}
