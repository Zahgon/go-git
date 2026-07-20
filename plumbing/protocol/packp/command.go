package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

type CommandArgs interface {
	Encoder
	Decoder
}

type CommandRequest struct {
	Command      string
	Capabilities capability.List
	Args         CommandArgs
}

func (c *CommandRequest) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (c *CommandRequest) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }
