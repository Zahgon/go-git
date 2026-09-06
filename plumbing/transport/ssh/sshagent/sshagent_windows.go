//go:build windows

package sshagent

import (
	"net"
	"sync"

	"golang.org/x/crypto/ssh/agent"
)

const (
	sshAgentPipe = `\\.\pipe\openssh-ssh-agent`
)

func Available() bool { _ = "STUB: not implemented"; return false }

func New() (agent.Agent, net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), *new(net.Conn), nil
}

type conn struct {
	sync.Mutex
	buf []byte
}

func (c *conn) Close() { _ = "STUB: not implemented"; return }

func (c *conn) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
