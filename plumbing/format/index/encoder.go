package index

import (
	"errors"
	"io"
	"time"

	"github.com/go-git/go-git/v6/plumbing/hash"
)

var (
	EncodeVersionSupported uint32 = 4

	ErrInvalidTimestamp = errors.New("negative timestamps are not allowed")
)

type Encoder struct {
	w         io.Writer
	hash      hash.Hash
	lastEntry *Entry
	skipHash  bool
}

func NewEncoder(w io.Writer, h hash.Hash, opts ...Option) *Encoder {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) Encode(idx *Index) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encode(idx *Index, footer bool) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeHeader(idx *Index) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeEntries(idx *Index) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeEntry(idx *Index, entry *Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeEntryName(entry *Entry) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeEntryNameV4(entry *Entry) error { _ = "STUB: not implemented"; return nil }

func commonPrefixLen(a, b string) int { _ = "STUB: not implemented"; return 0 }

func (e *Encoder) encodeRawExtension(signature string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) timeToUint32(t *time.Time) (uint32, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (e *Encoder) padEntry(idx *Index, wrote int) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeFooter() error { _ = "STUB: not implemented"; return nil }

type byName []*Entry

func (l byName) Len() int           { _ = "STUB: not implemented"; return 0 }
func (l byName) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (l byName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
