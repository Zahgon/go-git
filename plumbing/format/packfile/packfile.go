package packfile

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	billy "github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/cache"
	format "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
	"github.com/go-git/go-git/v6/plumbing/storer"
	gogitsync "github.com/go-git/go-git/v6/utils/sync"
)

var (
	ErrInvalidObject = NewError("invalid git object")

	ErrZLib = NewError("zlib reading error")
)

type Packfile struct {
	idxfile.Index
	fs   billy.Filesystem
	file billy.File

	handle        PackHandle
	resolveHandle PackHandleResolver

	scanReader io.ReadSeekCloser
	scanner    *Scanner

	cache cache.Object
	rbuf  *bufio.Reader

	id           plumbing.Hash
	m            sync.Mutex
	objectIDSize int

	once    sync.Once
	onceErr error

	closed atomic.Bool
}

func NewPackfile(
	file billy.File,
	opts ...PackfileOption,
) *Packfile {
	_ = "STUB: not implemented"
	return nil
}

func (p *Packfile) Get(h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (p *Packfile) GetByOffset(offset int64) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (p *Packfile) GetSizeByOffset(offset int64) (size int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Packfile) GetAll() (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func (p *Packfile) GetByType(typ plumbing.ObjectType) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func (p *Packfile) Scanner() (*Scanner, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Packfile) ID() (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (p *Packfile) get(h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (p *Packfile) getByOffset(offset int64) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (p *Packfile) init() error {
	p.once.Do(func() {
		if p.handle == nil && p.resolveHandle != nil {
			h, err := p.resolveHandle()
			if err != nil {
				p.onceErr = fmt.Errorf("packfile: resolve pack handle: %w", err)
				return
			}
			p.handle = h
		}

		if p.handle == nil && p.file == nil {
			p.onceErr = fmt.Errorf("file is not set")
			return
		}

		if p.Index == nil {
			p.onceErr = fmt.Errorf("index is not set")
			return
		}

		p.rbuf = gogitsync.GetBufioReader(nil)

		opts := []ScannerOption{WithBufioReader(p.rbuf)}

		if p.objectIDSize == format.SHA256Size {
			opts = append(opts, WithSHA256())
		}

		var scanSrc io.Reader
		if p.handle != nil {
			r, err := p.handle.OpenPackReader()
			if err != nil {
				p.onceErr = fmt.Errorf("packfile: open pack reader: %w", err)
				return
			}
			p.scanReader = r
			scanSrc = r
		} else {
			scanSrc = p.file
		}

		p.scanner = NewScanner(scanSrc, opts...)

		if !p.scanner.Scan() {
			p.onceErr = p.scanner.Error()
			return
		}

		if p.handle != nil {
			id, err := p.handle.PackHash()
			if err != nil {
				p.onceErr = fmt.Errorf("packfile: read pack hash: %w", err)
				return
			}
			p.id = id
		} else {
			_, err := p.scanner.Seek(-int64(p.objectIDSize), io.SeekEnd)
			if err != nil {
				p.onceErr = err
				return
			}
			p.id.ResetBySize(p.objectIDSize)
			_, err = p.id.ReadFrom(p.scanner)
			if err != nil {
				p.onceErr = err
			}
		}

		if p.cache == nil {
			p.cache = cache.NewObjectLRUDefault()
		}
	})

	return p.onceErr
}

func (p *Packfile) headerFromOffset(offset int64) (*ObjectHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Packfile) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Packfile) objectFromHeader(oh *ObjectHeader) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (p *Packfile) getMemoryObject(oh *ObjectHeader) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}
