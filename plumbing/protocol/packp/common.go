package packp

import (
	"fmt"
)

const (
	sha1HexSize   = 40
	sha256HexSize = 64
)

var (
	sp  = []byte(" ")
	eol = []byte("\n")

	null       = []byte("\x00")
	noHeadMark = []byte(" capabilities^{}\x00")

	want            = []byte("want ")
	shallow         = []byte("shallow ")
	deepen          = []byte("deepen")
	deepenCommits   = []byte("deepen ")
	deepenSince     = []byte("deepen-since ")
	deepenReference = []byte("deepen-not ")

	unshallow = []byte("unshallow ")

	ack = []byte("ACK")
	nak = []byte("NAK")

	shallowNoSp = []byte("shallow")
)

func isFlush(payload []byte) bool { _ = "STUB: not implemented"; return false }

var ErrNilWriter = fmt.Errorf("nil writer")

type ErrUnexpectedData struct {
	Msg  string
	Data []byte
}

func NewErrUnexpectedData(msg string, data []byte) error { _ = "STUB: not implemented"; return nil }

func (err *ErrUnexpectedData) Error() string { _ = "STUB: not implemented"; return "" }
