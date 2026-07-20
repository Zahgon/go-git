package packp

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

const ackLineLen = 44

type ServerResponse struct {
	ACKs []ACK
}

type ACKStatus byte

func (s ACKStatus) String() string { _ = "STUB: not implemented"; return "" }

const (
	ACKContinue ACKStatus = iota + 1
	ACKCommon
	ACKReady
)

type ACK struct {
	Hash   plumbing.Hash
	Status ACKStatus
}

func (r *ServerResponse) Decode(reader io.Reader) error { _ = "STUB: not implemented"; return nil }

func (r *ServerResponse) decodeLine(line []byte) error { _ = "STUB: not implemented"; return nil }

func (r *ServerResponse) decodeACKLine(line []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *ServerResponse) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func encodeServerResponse(w io.Writer, acks []ACK) error { _ = "STUB: not implemented"; return nil }
