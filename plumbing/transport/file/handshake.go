package file

import (
	"context"

	transport "github.com/go-git/go-git/v6/plumbing/transport"
)

func (t *Transport) Handshake(ctx context.Context, req *transport.Request) (transport.Session, error) {
	_ = "STUB: not implemented"
	return *new(transport.Session), nil
}

var (
	_ transport.Transport = (*Transport)(nil)
	_ transport.Connector = (*Transport)(nil)
)
