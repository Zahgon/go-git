package idxfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/go-git/go-git/v6/internal/sharedfile"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/x/fdpool"
)

const defaultCloseGracePeriod = time.Second

const (
	idxHeaderSize = 8
	idxFanoutSize = 256 * 4
	off32Size     = 4
	off64Size     = 8
	revHeaderSize = 12

	is64bitsMask = uint64(1) << 31
)

type ReadAtCloser = sharedfile.ReadAtCloser

type LazyIndex struct {
	hashSize int
	count    int
	count64  int

	fanoutStart int
	namesStart  int
	crcStart    int
	off32Start  int
	off64Start  int

	idx *sharedfile.SharedFile
	rev *sharedfile.SharedFile

	fanout [256]uint32
}

var _ Index = (*LazyIndex)(nil)

func NewLazyIndex(openIdx, openRev func() (ReadAtCloser, error), packHash plumbing.Hash) (*LazyIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLazyIndexWithPool(openIdx, openRev func() (ReadAtCloser, error), packHash plumbing.Hash, pool *fdpool.Pool) (*LazyIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *LazyIndex) init(packHash plumbing.Hash) error {
	idxRA, err := s.idx.Acquire()
	if err != nil {
		return fmt.Errorf("cannot open idx: %w", err)
	}
	defer s.idx.Release()

	revRA, err := s.rev.Acquire()
	if err != nil {
		return fmt.Errorf("cannot open rev: %w", err)
	}
	defer s.rev.Release()

	var hdr [idxHeaderSize]byte
	if _, err := idxRA.ReadAt(hdr[:], 0); err != nil {
		return fmt.Errorf("cannot read idx header: %w", err)
	}
	if !bytes.Equal(hdr[:4], idxHeader) {
		return fmt.Errorf("%w: %s", ErrMalformedIdxFile, "header mismatch")
	}

	v := binary.BigEndian.Uint32(hdr[4:])
	if v != VersionSupported {
		return ErrUnsupportedVersion
	}

	var revHdr [revHeaderSize]byte
	if _, err := revRA.ReadAt(revHdr[:], 0); err != nil {
		return fmt.Errorf("cannot read rev header: %w", err)
	}
	if !bytes.Equal(revHdr[:4], []byte{'R', 'I', 'D', 'X'}) {
		return fmt.Errorf("%w: rev file magic mismatch", ErrMalformedIdxFile)
	}
	if v := binary.BigEndian.Uint32(revHdr[4:]); v != 1 {
		return fmt.Errorf("%w: unsupported rev file version %d", ErrMalformedIdxFile, v)
	}

	s.fanoutStart = idxHeaderSize
	var fanoutBuf [idxFanoutSize]byte
	if _, err := idxRA.ReadAt(fanoutBuf[:], int64(s.fanoutStart)); err != nil {
		return fmt.Errorf("cannot read idx fanout: %w", err)
	}

	for i := range 256 {
		s.fanout[i] = binary.BigEndian.Uint32(fanoutBuf[i*4:])
		if i > 0 && s.fanout[i] < s.fanout[i-1] {
			return fmt.Errorf("%w: fanout table is not monotonically non-decreasing at entry %d",
				ErrMalformedIdxFile, i)
		}
	}

	s.count = int(s.fanout[255])

	s.hashSize = packHash.Size()
	s.namesStart = s.fanoutStart + idxFanoutSize
	s.crcStart = s.namesStart + (s.count * s.hashSize)
	s.off32Start = s.crcStart + (s.count * 4)
	s.off64Start = s.off32Start + (s.count * off32Size)

	n64, err := s.count64bitOffsets(idxRA)
	if err != nil {
		return err
	}
	s.count64 = n64

	packBuf := make([]byte, s.hashSize)
	packHashOff := int64(s.off64Start + n64*off64Size)
	if _, err := idxRA.ReadAt(packBuf, packHashOff); err != nil {
		return fmt.Errorf("cannot read pack checksum: %w", err)
	}

	if packHash.Compare(packBuf) != 0 {
		var got plumbing.Hash
		got.ResetBySize(s.hashSize)
		_, _ = got.Write(packBuf)
		return fmt.Errorf("%w: packfile mismatch: got %q instead of %q",
			ErrMalformedIdxFile, got.String(), packHash.String())
	}

	return nil
}

func (s *LazyIndex) Contains(h plumbing.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *LazyIndex) MayContain(h plumbing.Hash) bool { _ = "STUB: not implemented"; return false }

func (s *LazyIndex) FindOffset(h plumbing.Hash) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *LazyIndex) FindCRC32(h plumbing.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *LazyIndex) FindHash(o int64) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (s *LazyIndex) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *LazyIndex) Entries() (EntryIter, error) {
	_ = "STUB: not implemented"
	return *new(EntryIter), nil
}

func (s *LazyIndex) EntriesWithPrefix(prefix []byte) (EntryIter, error) {
	_ = "STUB: not implemented"
	return *new(EntryIter), nil
}

func (s *LazyIndex) EntriesByOffset() (EntryIter, error) {
	_ = "STUB: not implemented"
	return *new(EntryIter), nil
}

func (s *LazyIndex) Close() error { _ = "STUB: not implemented"; return nil }

func (s *LazyIndex) CloseIdleDescriptors() error { _ = "STUB: not implemented"; return nil }

func (s *LazyIndex) findHashPos(idx io.ReaderAt, h plumbing.Hash) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (s *LazyIndex) offset(idx io.ReaderAt, pos int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *LazyIndex) count64bitOffsets(idx io.ReaderAt) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *LazyIndex) crc32(idx io.ReaderAt, pos int) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *LazyIndex) hashAtPos(idx io.ReaderAt, pos int) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (s *LazyIndex) findHashViaRev(idx, rev io.ReaderAt, want int64) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (s *LazyIndex) entryAt(idx io.ReaderAt, pos int) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type scannerEntryIter struct {
	s   *LazyIndex
	idx io.ReaderAt
	pos int
}

func (it *scannerEntryIter) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (it *scannerEntryIter) Close() error { _ = "STUB: not implemented"; return nil }

type revEntryIter struct {
	s   *LazyIndex
	idx io.ReaderAt
	rev io.ReaderAt
	pos int
}

func (it *revEntryIter) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (it *revEntryIter) Close() error { _ = "STUB: not implemented"; return nil }

type lazyPrefixIter struct {
	s      *LazyIndex
	idx    io.ReaderAt
	prefix []byte
	pos    int
	end    int
}

func (it *lazyPrefixIter) Next() (*Entry, error) { _ = "STUB: not implemented"; return nil, nil }

func (it *lazyPrefixIter) Close() error { _ = "STUB: not implemented"; return nil }
