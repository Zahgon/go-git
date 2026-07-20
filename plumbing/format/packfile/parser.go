package packfile

import (
	"errors"
	"io"
	stdsync "sync"

	"github.com/go-git/go-git/v6/plumbing"
	format "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

var (
	ErrReferenceDeltaNotFound = errors.New("reference delta not found")

	ErrNotSeekableSource = errors.New("parser source is not seekable and storage was not provided")

	ErrDeltaNotCached = errors.New("delta could not be found in cache")

	ErrParserConsumed = errors.New("parser already consumed")
)

const maxObjectPreallocBytes = 1 << 30

const maxDeltaChainDepth = 4095

func growHint(n int64) int { _ = "STUB: not implemented"; return 0 }

type Parser struct {
	storage       storer.EncodedObjectStorer
	cache         *parserCache
	lowMemoryMode bool

	scanner   *Scanner
	observers []Observer
	hasher    plumbing.Hasher

	objectFormat format.ObjectFormat

	checksum plumbing.Hash
	m        stdsync.Mutex
	parsed   bool
}

type LowMemoryCapable interface {
	LowMemoryMode() bool
}

func NewParser(data io.Reader, opts ...ParserOption) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) storeOrCache(oh *ObjectHeader) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) resetCache(qty int) { _ = "STUB: not implemented"; return }

func (p *Parser) Parse() (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (p *Parser) ensureContent(oh *ObjectHeader) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) resolveDeltas(ofsDeltas, refDeltas []*ObjectHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) processDelta(oh *ObjectHeader) error { _ = "STUB: not implemented"; return nil }

func checkDeltaChainDepth(oh *ObjectHeader) error { _ = "STUB: not implemented"; return nil }

func (oh *ObjectHeader) isDeltaOnDisk() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) parentReader(parent *ObjectHeader) (io.ReaderAt, error) {
	_ = "STUB: not implemented"
	return *new(io.ReaderAt), nil
}

func (p *Parser) applyPatchBaseHeader(ota *ObjectHeader, delta io.Reader, target io.Writer, wh objectHeaderWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) forEachObserver(f func(o Observer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) onHeader(count uint32) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) onInflatedObjectHeader(
	t plumbing.ObjectType,
	objSize int64,
	pos int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) onInflatedObjectContent(
	h plumbing.Hash,
	pos int64,
	crc uint32,
	content []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) onFooter(h plumbing.Hash) error { _ = "STUB: not implemented"; return nil }
