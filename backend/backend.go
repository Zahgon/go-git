package backend

import (
	"context"
	"io"
	"log"
	"net/url"

	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/plumbing/transport"
)

type Request struct {
	URL *url.URL

	Service string

	GitProtocol string

	AdvertiseRefs bool

	StatelessRPC bool
}

type Backend struct {
	Loader transport.Loader

	ErrorLog *log.Logger

	Prefix string
}

func New(loader transport.Loader) *Backend { _ = "STUB: not implemented"; return nil }

func (b *Backend) Serve(ctx context.Context, r io.ReadCloser, w io.WriteCloser, req *Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) logf(format string, v ...any) { _ = "STUB: not implemented"; return }

func RequestFromProto(proto *packp.GitProtoRequest) *Request { _ = "STUB: not implemented"; return nil }
