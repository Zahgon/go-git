package ssh

import (
	"context"
	"net"

	gossh "golang.org/x/crypto/ssh"

	transport "github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh/knownhosts"
)

const (
	keyboardInteractiveName = "ssh-keyboard-interactive"
	passwordName            = "ssh-password"
	passwordCallbackName    = "ssh-password-callback"
	publicKeysName          = "ssh-public-keys"
	publicKeysCallbackName  = "ssh-public-key-callback"
)

type HostKeyCallbackHelper struct {
	HostKeyCallback gossh.HostKeyCallback
}

func (m *HostKeyCallbackHelper) SetHostKeyCallback(cfg *gossh.ClientConfig) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *HostKeyCallbackHelper) traceHostKeyCallback(hostname string, remote net.Addr, key gossh.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

type Password struct {
	User     string
	Password string
	HostKeyCallbackHelper
}

func (a *Password) ClientConfig(_ context.Context, _ *transport.Request) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PasswordCallback struct {
	User     string
	Callback func() (pass string, err error)
	HostKeyCallbackHelper
}

func (a *PasswordCallback) ClientConfig(_ context.Context, _ *transport.Request) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PublicKeys struct {
	User   string
	Signer gossh.Signer
	HostKeyCallbackHelper
}

func NewPublicKeys(user string, pemBytes []byte, password string) (*PublicKeys, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPublicKeysFromFile(user, pemFile, password string) (*PublicKeys, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *PublicKeys) ClientConfig(_ context.Context, _ *transport.Request) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PublicKeysCallback struct {
	User     string
	Callback func() (signers []gossh.Signer, err error)
	HostKeyCallbackHelper
}

func NewSSHAgentAuth(u string) (*PublicKeysCallback, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *PublicKeysCallback) ClientConfig(_ context.Context, _ *transport.Request) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type KeyboardInteractive struct {
	User      string
	Challenge gossh.KeyboardInteractiveChallenge
	HostKeyCallbackHelper
}

func (a *KeyboardInteractive) ClientConfig(_ context.Context, _ *transport.Request) (*gossh.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKnownHostsCallback(files ...string) (gossh.HostKeyCallback, error) {
	_ = "STUB: not implemented"
	return *new(gossh.HostKeyCallback), nil
}

func newKnownHostsDb(files ...string) (*knownhosts.HostKeyDB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDefaultKnownHostsFiles() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func tracePublicKeysCallback(getSigners func() ([]gossh.Signer, error)) gossh.AuthMethod {
	_ = "STUB: not implemented"
	return *new(gossh.AuthMethod)
}

func filterKnownHostsFiles(files ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func username() (string, error) { _ = "STUB: not implemented"; return "", nil }
