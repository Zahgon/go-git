package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	transport "github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/storage"
)

func (t *Transport) Handshake(ctx context.Context, req *transport.Request) (transport.Session, error) {
	_ = "STUB: not implemented"
	return *new(transport.Session), nil
}

func handshakeSmart(resp *http.Response, req *transport.Request, discoverService string, client *http.Client, authorizer func(*http.Request) error) (transport.Session, error) {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck
	new(transport.Session), nil
}

func handshakeDumb(resp *http.Response, req *transport.Request, client *http.Client, authorizer func(*http.Request) error) (transport.Session, error) {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck
	new(transport.Session), nil
}

var (
	_ transport.Session   = (*smartPackSession)(nil)
	_ transport.Commander = (*smartPackSession)(nil)
	_ transport.Archiver  = (*smartPackSession)(nil)
)

type smartPackSession struct {
	client     *http.Client
	baseURL    *url.URL
	service    string
	authorizer func(*http.Request) error
	version    protocol.Version
	caps       capability.List
	refs       *packp.AdvRefs
}

func (s *smartPackSession) Capabilities() *capability.List { _ = "STUB: not implemented"; return nil }

func (s *smartPackSession) GetRemoteRefs(ctx context.Context, opts *transport.GetRemoteRefsOptions) (*transport.RemoteRefs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *smartPackSession) Command(ctx context.Context, cmd string, req packp.CommandArgs, resp packp.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *smartPackSession) Fetch(ctx context.Context, st storage.Storer, req *transport.FetchRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *smartPackSession) fetchV2(ctx context.Context, st storage.Storer, req *transport.FetchRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *smartPackSession) Push(ctx context.Context, st storage.Storer, req *transport.PushRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *smartPackSession) Close() error { _ = "STUB: not implemented"; return nil }

func (s *smartPackSession) Archive(ctx context.Context, req *transport.ArchiveRequest) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type httpArchiveBody struct{ req *httpRequester }

func (b *httpArchiveBody) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *httpArchiveBody) Close() error { _ = "STUB: not implemented"; return nil }

type httpRequester struct {
	session *smartPackSession
	ctx     context.Context
	buf     bytes.Buffer
	resp    *http.Response
}

func (r *httpRequester) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *httpRequester) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *httpRequester) Close() error { _ = "STUB: not implemented"; return nil }

func (r *httpRequester) doPost() error { _ = "STUB: not implemented"; return nil }

type httpNegotiator struct {
	session *smartPackSession
	ctx     context.Context
	current *httpRequester
}

func (n *httpNegotiator) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *httpNegotiator) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *httpNegotiator) Close() error { _ = "STUB: not implemented"; return nil }

func (n *httpNegotiator) closeResponse() { _ = "STUB: not implemented"; return }

var _ transport.Session = (*dumbPackSession)(nil)

type dumbPackSession struct {
	client     *http.Client
	baseURL    *url.URL
	service    string
	authorizer func(*http.Request) error
	refs       *packp.AdvRefs
}

func (s *dumbPackSession) Capabilities() *capability.List { _ = "STUB: not implemented"; return nil }

func (s *dumbPackSession) GetRemoteRefs(_ context.Context, _ *transport.GetRemoteRefsOptions) (*transport.RemoteRefs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dumbPackSession) Fetch(ctx context.Context, st storage.Storer, req *transport.FetchRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *dumbPackSession) Push(_ context.Context, _ storage.Storer, _ *transport.PushRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *dumbPackSession) Close() error { _ = "STUB: not implemented"; return nil }

var (
	_ transport.Session   = (*smartPackSession)(nil)
	_ transport.Session   = (*dumbPackSession)(nil)
	_ transport.Transport = (*Transport)(nil)
)
