package memory

import (
	"fmt"
	"io"
	"time"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	formatcfg "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/format/index"
	"github.com/go-git/go-git/v6/plumbing/format/reflog"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage"
)

var ErrUnsupportedObjectType = fmt.Errorf("unsupported object type")

type Storage struct {
	ConfigStorage
	ObjectStorage
	ShallowStorage
	IndexStorage
	ReferenceStorage
	ModuleStorage
	ReflogStorage
	options options
}

func NewStorage(o ...StorageOption) *Storage { _ = "STUB: not implemented"; return nil }

func (s *Storage) SetObjectFormat(of formatcfg.ObjectFormat) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Storage) SupportsExtension(name, value string) bool {
	_ = "STUB: not implemented"
	return false
}

type ConfigStorage struct {
	config *config.Config
}

func (c *ConfigStorage) SetConfig(cfg *config.Config) error { _ = "STUB: not implemented"; return nil }

func (c *ConfigStorage) Config() (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type IndexStorage struct {
	index *index.Index
}

func (c *IndexStorage) SetIndex(idx *index.Index) error { _ = "STUB: not implemented"; return nil }

func (c *IndexStorage) Index() (*index.Index, error) { _ = "STUB: not implemented"; return nil, nil }

type ObjectStorage struct {
	oh      *plumbing.ObjectHasher
	Objects map[plumbing.Hash]plumbing.EncodedObject
	Commits map[plumbing.Hash]plumbing.EncodedObject
	Trees   map[plumbing.Hash]plumbing.EncodedObject
	Blobs   map[plumbing.Hash]plumbing.EncodedObject
	Tags    map[plumbing.Hash]plumbing.EncodedObject
}

type lazyCloser struct {
	storage *ObjectStorage
	obj     plumbing.EncodedObject
	closer  io.Closer
}

func (c *lazyCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (o *ObjectStorage) RawObjectWriter(typ plumbing.ObjectType, sz int64) (w io.WriteCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (o *ObjectStorage) NewEncodedObject() plumbing.EncodedObject {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject)
}

func (o *ObjectStorage) SetEncodedObject(obj plumbing.EncodedObject) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (o *ObjectStorage) HasEncodedObject(h plumbing.Hash) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectStorage) EncodedObjectSize(h plumbing.Hash) (
	size int64, err error,
) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ObjectStorage) EncodedObject(t plumbing.ObjectType, h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (o *ObjectStorage) IterEncodedObjects(t plumbing.ObjectType) (storer.EncodedObjectIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.EncodedObjectIter), nil
}

func flattenObjectMap(m map[plumbing.Hash]plumbing.EncodedObject) []plumbing.EncodedObject {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectStorage) Begin() storer.Transaction {
	_ = "STUB: not implemented"
	return *new(storer.Transaction)
}

func (o *ObjectStorage) ForEachObjectHash(fun func(plumbing.Hash) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectStorage) ObjectPacks() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *ObjectStorage) DeleteOldObjectPackAndIndex(plumbing.Hash, time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

var errNotSupported = fmt.Errorf("not supported")

func (o *ObjectStorage) LooseObjectTime(_ plumbing.Hash) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (o *ObjectStorage) DeleteLooseObject(plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectStorage) AddAlternate(_ string) error { _ = "STUB: not implemented"; return nil }

type TxObjectStorage struct {
	Storage *ObjectStorage
	Objects map[plumbing.Hash]plumbing.EncodedObject
}

func (tx *TxObjectStorage) SetEncodedObject(obj plumbing.EncodedObject) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (tx *TxObjectStorage) EncodedObject(t plumbing.ObjectType, h plumbing.Hash) (plumbing.EncodedObject, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.EncodedObject), nil
}

func (tx *TxObjectStorage) Commit() error { _ = "STUB: not implemented"; return nil }

func (tx *TxObjectStorage) Rollback() error { _ = "STUB: not implemented"; return nil }

type ReferenceStorage map[plumbing.ReferenceName]*plumbing.Reference

func (r ReferenceStorage) SetReference(ref *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (r ReferenceStorage) CheckAndSetReference(ref, old *plumbing.Reference) error {
	_ = "STUB: not implemented"
	return nil
}

func (r ReferenceStorage) Reference(n plumbing.ReferenceName) (*plumbing.Reference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReferenceStorage) IterReferences() (storer.ReferenceIter, error) {
	_ = "STUB: not implemented"
	return *new(storer.ReferenceIter), nil
}

func (r ReferenceStorage) CountLooseRefs() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r ReferenceStorage) PackRefs() error { _ = "STUB: not implemented"; return nil }

func (r ReferenceStorage) RemoveReference(n plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

type ShallowStorage []plumbing.Hash

func (s *ShallowStorage) SetShallow(commits []plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ShallowStorage) Shallow() ([]plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ModuleStorage map[string]*Storage

func (s ModuleStorage) Module(name string) (storage.Storer, error) {
	_ = "STUB: not implemented"
	return *new(storage.Storer), nil
}

type ReflogStorage struct {
	entries map[plumbing.ReferenceName][]*reflog.Entry
}

func (r *ReflogStorage) Reflog(name plumbing.ReferenceName) ([]*reflog.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReflogStorage) AppendReflog(name plumbing.ReferenceName, entry *reflog.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReflogStorage) DeleteReflog(name plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}
