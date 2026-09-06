package packp

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing/protocol"
	"github.com/go-git/go-git/v6/plumbing/protocol/capability"
)

type CapabilityAdv struct {
	Version protocol.Version

	Capabilities capability.List
}

func (ca *CapabilityAdv) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (ca *CapabilityAdv) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

var errInvalidVersionLine = errors.New("capability advertisement must start with version line")
