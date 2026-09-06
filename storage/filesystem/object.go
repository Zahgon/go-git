package filesystem

import (
	"io"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/plumbing/format/idxfile"
	"github.com/go-git/go-git/v6/plumbing/format/packfile"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage/filesystem/dotgit"
)

const indexSFKey = "populate"

const reindexSFKey = "reindex"

type packEntry struct {
	h   plumbing.Hash
	idx idxfile.Index
}

type ObjectStorage struct {
	options Options

	objectCache cache.Object

	dir   *dotgit.DotGit
	index map[plumbing.Hash]idxfile.Index

	packs []packEntry
	muI   sync.RWMutex

	indexSF singleflight.Group

	lastHitPackIdx atomic.Int32

	oh *plumbing.ObjectHasher

	alternates     []*ObjectStorage
	alternatesInit bool
	alternatesErr  error
	muA            sync.RWMutex
}

func NewObjectStorage(dir *dotgit.DotGit, objectCache cache.Object) *ObjectStorage {
	_ = "STUB: not implemented"
	return nil
}

func NewObjectStorageWithOptions(dir *dotgit.DotGit, objectCache cache.Object, ops Options) *ObjectStorage {
	_ = "STUB: not implemented"
	return nil
}

func (s *ObjectStorage) initAlternates() error { _ = "STUB: not implemented"; return nil }

func (s *ObjectStorage) resetAlternates() { _ = "STUB: not implemented"; return }

func findInAlternates[T any](s *ObjectStorage, fn func(*ObjectStorage) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (s *ObjectStorage) requireIndex() error { _ = "STUB: not implemented"; return nil }

func (s *ObjectStorage) Reindex() error { _ = "STUB: not implemented"; return nil }

func (s *ObjectStorage) populateIndex() (map[plumbing.Hash]idxfile.Index, []packEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ObjectStorage) loadIdx(h plumbing.Hash) (idxfile.Index, error) {
	_ = "STUB: not implemented"
	return *new(idxfile.Index), nil
}

func (s *ObjectStorage) loadLazyIndex(h plumbing.Hash) (*idxfile.LazyIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ObjectStorage) loadMemoryIndexValue(h plumbing.Hash) (idx idxfile.Index, err error) {
	_ = "STUB: not implemented"
	return *new(idxfile.Index), nil
}

func (s *ObjectStorage) RawObjectWriter(typ plumbing.ObjectType, sz int64) (w io.WriteCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (s *ObjectStorage) NewEncodedObject() plumbing.EncodedObject {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject)
}

func (s *ObjectStorage) PackfileWriter() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (s *ObjectStorage) PromisorPackfileWriter(marker string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (s *ObjectStorage) PromisorObjectPacks() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ObjectStorage) packfileWriter(newPack func() (*dotgit.PackWriter, error)) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (s *ObjectStorage) SetEncodedObject(o plumbing.EncodedObject) (h plumbing.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (s *ObjectStorage) LazyWriter() (w io.WriteCloser, wh func(typ plumbing.ObjectType, sz int64) error, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil, nil
}

func (s *ObjectStorage) HasEncodedObject(h plumbing.Hash) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *ObjectStorage) encodedObjectSizeFromUnpacked(h plumbing.Hash) (size int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *ObjectStorage) packfile(idx idxfile.Index, pack plumbing.Hash) (*packfile.Packfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ObjectStorage) EncodedObjectSize(h plumbing.Hash) (size int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *ObjectStorage) EncodedObject(t plumbing.ObjectType, h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *ObjectStorage) DeltaObject(t plumbing.ObjectType, h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *ObjectStorage) getFromUnpacked(h plumbing.Hash) (obj plumbing.EncodedObject, err error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *ObjectStorage) getFromPackfile(h plumbing.Hash, canBeDelta bool) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *ObjectStorage) getFromPackfileAt(pack plumbing.Hash, idx idxfile.Index, h plumbing.Hash, offset int64, canBeDelta bool) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (s *ObjectStorage) decodeDeltaObjectAt(
	p *packfile.Packfile,
	offset int64,
	hash plumbing.Hash,
) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *
	//nolint:staticcheck // TODO: Refactor to avoid deprecated Scanner method
	new(plumbing.EncodedObject), nil
}

func (s *ObjectStorage) findObjectInPackfile(h plumbing.Hash) (plumbing.Hash, idxfile.Index, int64) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), *new(idxfile.Index), 0
}

func (s *ObjectStorage) HashesWithPrefix(prefix []byte) ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ObjectStorage) IterEncodedObjects(t plumbing.ObjectType) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func (s *ObjectStorage) buildPackfileIters(
	t plumbing.ObjectType,
	seen map[plumbing.Hash]struct{},
) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func (s *ObjectStorage) Close() error { _ = "STUB: not implemented"; return nil }

func (s *ObjectStorage) CloseIdleDescriptors() error { _ = "STUB: not implemented"; return nil }

func hashListAsMap(l []plumbing.Hash) map[plumbing.Hash]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *ObjectStorage) ForEachObjectHash(fun func(plumbing.Hash) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ObjectStorage) LooseObjectTime(hash plumbing.Hash) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (s *ObjectStorage) DeleteLooseObject(hash plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ObjectStorage) ObjectPacks() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ObjectStorage) DeleteOldObjectPackAndIndex(h plumbing.Hash, t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
