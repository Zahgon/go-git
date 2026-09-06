package plumbing

import (
	"io"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
)

var empty = make([]byte, format.SHA256Size)

func FromHex(in string) (ObjectID, bool) { _ = "STUB: not implemented"; return *new(ObjectID), false }

func FromBytes(in []byte) (ObjectID, bool) { _ = "STUB: not implemented"; return *new(ObjectID), false }

type ObjectID struct {
	hash   [format.SHA256Size]byte
	format format.ObjectFormat
}

func (s ObjectID) HexSize() int { _ = "STUB: not implemented"; return 0 }

func (s ObjectID) Size() int { _ = "STUB: not implemented"; return 0 }

func (s ObjectID) Compare(b []byte) int { _ = "STUB: not implemented"; return 0 }

func (s ObjectID) Equal(in ObjectID) bool { _ = "STUB: not implemented"; return false }

func (s ObjectID) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (s ObjectID) HasPrefix(prefix []byte) bool { _ = "STUB: not implemented"; return false }

func (s ObjectID) IsZero() bool { _ = "STUB: not implemented"; return false }

func (s ObjectID) String() string { _ = "STUB: not implemented"; return "" }

func (s *ObjectID) Write(in []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ObjectID) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ObjectID) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ObjectID) ResetBySize(idSize int) { _ = "STUB: not implemented"; return }
