//go:build !windows

package sshagent

import (
	"net"

	"golang.org/x/crypto/ssh/agent"
)

func New() (agent.Agent, net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), *new(net.Conn), nil
}

func Available() bool { _ = "STUB: not implemented"; return false }
