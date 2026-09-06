package packp

import (
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

var (
	minCommandLength        = sha1HexSize*2 + 2 + 1
	minCommandAndCapsLength = minCommandLength + 1
)

var (
	ErrEmpty                        = errors.New("empty update-request message")
	errNoCommands                   = errors.New("unexpected EOF before any command")
	errMissingCapabilitiesDelimiter = errors.New("capabilities delimiter not found")
	errNoFlush                      = errors.New("unexpected EOF before flush line")
)

func errMalformedRequest(reason string) error { _ = "STUB: not implemented"; return nil }

func errInvalidHash(hash string) error { _ = "STUB: not implemented"; return nil }

func errInvalidShallowLineLength(got int) error { _ = "STUB: not implemented"; return nil }

func errInvalidCommandCapabilitiesLineLength(got int) error { _ = "STUB: not implemented"; return nil }

func errInvalidCommandLineLength(got int) error { _ = "STUB: not implemented"; return nil }

func errInvalidShallowObjID(err error) error { _ = "STUB: not implemented"; return nil }

func errInvalidOldObjID(err error) error { _ = "STUB: not implemented"; return nil }

func errInvalidNewObjID(err error) error { _ = "STUB: not implemented"; return nil }

func errMalformedCommand(err error) error { _ = "STUB: not implemented"; return nil }

func (req *UpdateRequests) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func parseCommand(b []byte) (*Command, error) { _ = "STUB: not implemented"; return nil, nil }

func parseHash(s string) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}
