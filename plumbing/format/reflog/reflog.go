package reflog

import (
	"bufio"
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing"
)

type Signature struct {
	Name string

	Email string

	When time.Time
}

type Entry struct {
	OldHash plumbing.Hash

	NewHash plumbing.Hash

	Committer Signature

	Message string
}

type Decoder struct {
	r *bufio.Reader
}

func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

func (d *Decoder) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func Decode(r io.Reader) ([]*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func decodeLine(line []byte) (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func decodeTimestamp(s []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func normalizeMessage(msg string) string { _ = "STUB: not implemented"; return "" }

func Encode(w io.Writer, e *Entry) error { _ = "STUB: not implemented"; return nil }
