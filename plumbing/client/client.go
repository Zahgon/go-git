package client

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	gossh "golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"

	"github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/plumbing/transport/file"
	xgit "github.com/go-git/go-git/v6/plumbing/transport/git"
	xhttp "github.com/go-git/go-git/v6/plumbing/transport/http"
	xssh "github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

type SSHAuth interface {
	ClientConfig(context.Context, *transport.Request) (*gossh.ClientConfig, error)
}

type HTTPAuth interface {
	Authorizer(*http.Request) error
}

type RedirectPolicy = xhttp.RedirectPolicy

const (
	FollowInitialRedirects = xhttp.FollowInitialRedirects

	FollowRedirects = xhttp.FollowRedirects

	NoFollowRedirects = xhttp.NoFollowRedirects
)

type Option func(*options)

type options struct {
	ssh  xssh.Options
	http xhttp.Options
	git  xgit.Options
	file file.Options

	schemes map[string]transport.Transport
}

func (o *options) ensureTLS() *tls.Config { _ = "STUB: not implemented"; return nil }

func WithSSHAuth(a SSHAuth) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPAuth(a HTTPAuth) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProxyURL(u *url.URL) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProxyEnvironment() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialer(fn transport.DialContextFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHTTPClient(c *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRedirectPolicy(policy RedirectPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithInsecureSkipTLS() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCABundle(pem []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLoader(l transport.Loader) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTransport(scheme string, tr transport.Transport) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type Client struct {
	opts options
}

func New(opts ...Option) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Handshake(ctx context.Context, req *transport.Request) (transport.Session, error) {
	_ = "STUB: not implemented"
	return *new(transport.Session), nil
}

func (c *Client) Connect(ctx context.Context, req *transport.Request) (transport.Conn, error) {
	_ = "STUB: not implemented"
	return *new(transport.Conn), nil
}

func (c *Client) Transport(scheme string) (transport.Transport, error) {
	_ = "STUB: not implemented"
	return *new(transport.Transport), nil
}

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Client) resolve(req *transport.Request) (transport.Transport, error) {
	_ = "STUB: not implemented"
	return *new(transport.Transport), nil
}

func (c *Client) builtin(scheme string) (transport.Transport, error) {
	_ = "STUB: not implemented"
	return *new(transport.Transport), nil
}

func proxyDialer(makeDialer func(proxy.Dialer) (proxy.Dialer, error)) func(transport.DialContextFunc) transport.DialContextFunc {
	_ = "STUB: not implemented"
	return nil
}
