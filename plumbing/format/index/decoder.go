package index

import (
	"bufio"
	"errors"
	"io"

	"github.com/go-git/go-git/v6/plumbing/hash"
)

var (
	DecodeVersionSupported = struct{ Min, Max uint32 }{Min: 2, Max: 4}

	ErrMalformedSignature = errors.New("index decoder: malformed index signature file")

	ErrInvalidChecksum = errors.New("index decoder: invalid checksum")

	ErrUnknownExtension = errors.New("index decoder: unknown extension")

	ErrMalformedIndexFile = errors.New("index decoder: malformed index file")
)

const (
	entryHeaderLength = 42
	entryExtended     = 0x4000
	nameMask          = 0xfff
	intentToAddMask   = 1 << 13
	skipWorkTreeMask  = 1 << 14
)

type Decoder struct {
	buf       *bufio.Reader
	r         io.Reader
	hash      hash.Hash
	lastEntry *Entry
	skipHash  bool

	extReader *bufio.Reader
}

func NewDecoder(r io.Reader, h hash.Hash, opts ...Option) *Decoder {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) Decode(idx *Index) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) readEntries(idx *Index, count int) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) readEntry(idx *Index) (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Decoder) readEntryName(idx *Index, e *Entry, flags uint16) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Decoder) doReadEntryName(nameLen uint16) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (d *Decoder) doReadEntryNameV4() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *Decoder) padEntry(idx *Index, e *Entry, read, nameConsumed int) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Decoder) readExtensions(idx *Index) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) readExtension(idx *Index) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) getExtensionReader() (*bufio.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Decoder) readChecksum(expected []byte) error { _ = "STUB: not implemented"; return nil }

func validateHeader(r io.Reader) (version uint32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type treeExtensionDecoder struct {
	r *bufio.Reader
	h hash.Hash
}

func (d *treeExtensionDecoder) Decode(t *Tree) error { _ = "STUB: not implemented"; return nil }

func (d *treeExtensionDecoder) readEntry() (*TreeEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type resolveUndoDecoder struct {
	r *bufio.Reader
	h hash.Hash
}

func (d *resolveUndoDecoder) Decode(ru *ResolveUndo) error { _ = "STUB: not implemented"; return nil }

func (d *resolveUndoDecoder) readEntry() (*ResolveUndoEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *resolveUndoDecoder) readStage(e *ResolveUndoEntry, s Stage) error {
	_ = "STUB: not implemented"
	return nil
}

type endOfIndexEntryDecoder struct {
	r *bufio.Reader
	h hash.Hash
}

func (d *endOfIndexEntryDecoder) Decode(e *EndOfIndexEntry) error {
	_ = "STUB: not implemented"
	return nil
}

type unknownExtensionDecoder struct {
	r *bufio.Reader
}

func (d *unknownExtensionDecoder) Decode() error { _ = "STUB: not implemented"; return nil }
