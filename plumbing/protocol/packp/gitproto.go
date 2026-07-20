package packp

import (
	"fmt"
	"io"
)

var ErrInvalidGitProtoRequest = fmt.Errorf("invalid git protocol request")

type GitProtoRequest struct {
	RequestCommand string
	Pathname       string

	Host string

	ExtraParams []string
}

func (g *GitProtoRequest) validate() error { _ = "STUB: not implemented"; return nil }

func (g *GitProtoRequest) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (g *GitProtoRequest) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }
