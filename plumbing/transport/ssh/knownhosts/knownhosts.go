package knownhosts

import (
	"io"
	"net"

	"golang.org/x/crypto/ssh"
)

type HostKeyDB struct {
	callback   ssh.HostKeyCallback
	isCert     map[string]bool
	isWildcard map[string]bool
}

func NewDB(files ...string) (*HostKeyDB, error) { _ = "STUB: not implemented"; return nil, nil }

func (hkdb *HostKeyDB) HostKeyCallback() ssh.HostKeyCallback {
	_ = "STUB: not implemented"
	return *new(ssh.HostKeyCallback)
}

type PublicKey struct {
	ssh.PublicKey
	Cert bool
}

func (hkdb *HostKeyDB) HostKeys(hostWithPort string) (keys []PublicKey) {
	_ = "STUB: not implemented"
	return nil
}

func (hkdb *HostKeyDB) HostKeyAlgorithms(hostWithPort string) (algos []string) {
	_ = "STUB: not implemented"
	return nil
}

func keyTypeToCertAlgo(keyType string) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck // DSA support needed for legacy SSH servers
//nolint:staticcheck // DSA support needed for legacy SSH servers

type HostKeyCallback ssh.HostKeyCallback

func New(files ...string) (HostKeyCallback, error) {
	_ = "STUB: not implemented"
	return *new(HostKeyCallback), nil
}

func (hkcb HostKeyCallback) HostKeyCallback() ssh.HostKeyCallback {
	_ = "STUB: not implemented"
	return *new(ssh.HostKeyCallback)
}

func (hkcb HostKeyCallback) ToDB() *HostKeyDB { _ = "STUB: not implemented"; return nil }

func (hkcb HostKeyCallback) HostKeys(hostWithPort string) []ssh.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func (hkcb HostKeyCallback) HostKeyAlgorithms(hostWithPort string) (algos []string) {
	_ = "STUB: not implemented"
	return nil
}

func HostKeyAlgorithms(cb ssh.HostKeyCallback, hostWithPort string) []string {
	_ = "STUB: not implemented"
	return nil
}

func IsHostKeyChanged(err error) bool { _ = "STUB: not implemented"; return false }

func IsHostUnknown(err error) bool { _ = "STUB: not implemented"; return false }

func Normalize(address string) string { _ = "STUB: not implemented"; return "" }

func Line(addresses []string, key ssh.PublicKey) string { _ = "STUB: not implemented"; return "" }

func WriteKnownHost(w io.Writer, hostname string, remote net.Addr, key ssh.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteKnownHostCA(w io.Writer, hostPattern string, key ssh.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

type fakePublicKey struct{}

func (fakePublicKey) Type() string { _ = "STUB: not implemented"; return "" }

func (fakePublicKey) Marshal() []byte { _ = "STUB: not implemented"; return nil }

func (fakePublicKey) Verify(_ []byte, _ *ssh.Signature) error {
	_ = "STUB: not implemented"
	return nil
}
