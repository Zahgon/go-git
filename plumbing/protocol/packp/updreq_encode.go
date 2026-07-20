package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

func (req *UpdateRequests) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (req *UpdateRequests) encodeShallow(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (req *UpdateRequests) encodeCommands(w io.Writer,
	cmds []*Command, caps *capability.List,
) error {
	_ = "STUB: not implemented"
	return nil
}

func formatCommand(cmd *Command) string { _ = "STUB: not implemented"; return "" }
