package packp

import (
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

var (
	ErrEmptyCommands    = errors.New("commands cannot be empty")
	ErrMalformedCommand = errors.New("malformed command")
)

type UpdateRequests struct {
	Capabilities capability.List
	Commands     []*Command
	Shallows     []plumbing.Hash
}

func validateUpdateRequests(req *UpdateRequests) error { _ = "STUB: not implemented"; return nil }

type Action string

const (
	Create  Action = "create"
	Update  Action = "update"
	Delete  Action = "delete"
	Invalid Action = "invalid"
)

type Command struct {
	Name plumbing.ReferenceName
	Old  plumbing.Hash
	New  plumbing.Hash
}

func (c *Command) Action() Action { _ = "STUB: not implemented"; return *new(Action) }

func (c *Command) validate() error { _ = "STUB: not implemented"; return nil }
