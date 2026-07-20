package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/go-git/go-git/v6/plumbing/transport"
)

type contextKey int

const initialRequestKey contextKey = iota

type RedirectPolicy string

const (
	FollowInitialRedirects RedirectPolicy = "initial"

	FollowRedirects RedirectPolicy = "true"

	NoFollowRedirects RedirectPolicy = "false"
)

func withInitialRequest(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func isInitialRequest(req *http.Request) bool { _ = "STUB: not implemented"; return false }

type Options struct {
	Client *http.Client

	FollowRedirects RedirectPolicy

	Authorizer func(*http.Request) error

	HTTPProxy func(*http.Request) (*url.URL, error)

	TLS *tls.Config

	ForceDumb bool
}

type Transport struct {
	opts Options
}

var _ transport.Transport = (*Transport)(nil)

func NewTransport(opts Options) *Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) resolveClient() *http.Client { _ = "STUB: not implemented"; return nil }

func (o Options) redirectPolicy() RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

func wrapCheckRedirect(policy RedirectPolicy, next func(*http.Request, []*http.Request) error) func(*http.Request, []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func checkRedirect(req *http.Request, via []*http.Request, policy RedirectPolicy) error {
	_ = "STUB: not implemented"
	return nil
}
