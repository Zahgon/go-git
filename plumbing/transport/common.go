package transport

import (
	"context"
	"io"
	"net"

	internal "github.com/go-git/go-git/v6/internal/transport"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp/sideband"
)

type DialContextFunc func(ctx context.Context, network, address string) (net.Conn, error)

func (f DialContextFunc) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (f DialContextFunc) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type RemoteError struct {
	Reason string
}

func (e *RemoteError) Error() string { _ = "STUB: not implemented"; return "" }

func NewRemoteError(reason string) error { _ = "STUB: not implemented"; return nil }

type FetchRequest = internal.FetchRequest

type PushRequest struct {
	Packfile io.ReadCloser

	Commands []*packp.Command

	Progress sideband.Progress

	Options []string

	Atomic bool

	Quiet bool
}
