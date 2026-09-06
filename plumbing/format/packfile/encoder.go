package packfile

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/hash"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/utils/sync"
)

type ObjectSelector interface {
	ObjectsToPack(hashes []plumbing.Hash, packWindow uint) ([]*ObjectToPack, error)
}

type Encoder struct {
	deltaSelector  *DeltaSelector
	objectSelector ObjectSelector
	w              *offsetWriter
	zw             sync.ZlibWriter
	hasher         hash.Hash

	useRefDeltas bool
}

type EncoderOption func(*Encoder)

func WithObjectSelector(s ObjectSelector) EncoderOption {
	_ = "STUB: not implemented"
	return *new(EncoderOption)
}

func NewEncoder(w io.Writer, s storer.EncodedObjectStorer, useRefDeltas bool, opts ...EncoderOption) *Encoder {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) Encode(
	hashes []plumbing.Hash,
	packWindow uint,
) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (e *Encoder) encode(objects []*ObjectToPack) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (e *Encoder) head(numEntries int) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) entry(o *ObjectToPack) (err error) { _ = "STUB: not implemented"; return nil }

func (e *Encoder) writeBaseIfDelta(o *ObjectToPack) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) writeDeltaHeader(o *ObjectToPack) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) writeRefDeltaHeader(base plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) writeOfsDeltaHeader(o *ObjectToPack) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) entryHead(typeNum plumbing.ObjectType, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) footer() (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

type offsetWriter struct {
	w      io.Writer
	offset int64
}

func newOffsetWriter(w io.Writer) *offsetWriter { _ = "STUB: not implemented"; return nil }

func (ow *offsetWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ow *offsetWriter) Offset() int64 { _ = "STUB: not implemented"; return 0 }
