package ssh

import (
	"bytes"
	"context"
	"io"
	"strings"
	"sync/atomic"

	"github.com/kevinburke/ssh_config"
	gossh "golang.org/x/crypto/ssh"

	"github.com/go-git/go-git/v6/plumbing/transport"
)

const DefaultPort = 22

const DefaultUsername = "git"

type Options struct {
	ClientConfig func(context.Context, *transport.Request) (*gossh.ClientConfig, error)

	DialContext transport.DialContextFunc

	DialProxy func(transport.DialContextFunc) transport.DialContextFunc

	UserSettings func(context.Context, *transport.Request) (*ssh_config.UserSettings, error)
}

type Transport struct {
	opts Options
}

func NewTransport(opts Options) *Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) Connect(ctx context.Context, req *transport.Request) (transport.Conn, error) {
	_ = "STUB: not implemented"
	return *new(transport.Conn), nil
}

func (t *Transport) connect(ctx context.Context, req *transport.Request) (*sshConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) resolveConfig(ctx context.Context, req *transport.Request) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) dial(ctx context.Context, network, addr string, config *gossh.ClientConfig) (*gossh.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) userSettings(ctx context.Context, req *transport.Request) (*ssh_config.UserSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) resolveHostWithPort(ctx context.Context, req *transport.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type sshConn struct {
	stdout    io.Reader
	stdin     io.WriteCloser
	session   *gossh.Session
	client    *gossh.Client
	stderrBuf atomic.Pointer[bytes.Buffer]
}

var _ transport.Conn = (*sshConn)(nil)

func (c *sshConn) Reader() io.Reader      { _ = "STUB: not implemented"; return *new(io.Reader) }
func (c *sshConn) Writer() io.WriteCloser { _ = "STUB: not implemented"; return *new(io.WriteCloser) }

func (c *sshConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *sshConn) Stderr() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func buildCommand(req *transport.Request) string { _ = "STUB: not implemented"; return "" }

func writeShellQuote(b *strings.Builder, s string) { _ = "STUB: not implemented"; return }
