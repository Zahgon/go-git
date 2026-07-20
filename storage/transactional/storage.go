package transactional

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/format/reflog"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/storage"
)

type Storage interface {
	storage.Storer
	Commit() error
}

type basic struct {
	s, temporal storage.Storer

	*ObjectStorage
	*ReferenceStorage
	*IndexStorage
	*ShallowStorage
	*ConfigStorage
	reflog *ReflogStorage
}

type packageWriter struct {
	*basic
	pw storer.PackfileWriter
}

type reflogBasic struct {
	*basic
}

type reflogPackageWriter struct {
	*reflogBasic
	pw storer.PackfileWriter
}

func NewStorage(base, temporal storage.Storer) Storage {
	_ = "STUB: not implemented"
	return *new(Storage)
}

func (s *basic) Module(name string) (storage.Storer, error) {
	_ = "STUB: not implemented"
	return *new(storage.Storer), nil
}

func (s *basic) Commit() error { _ = "STUB: not implemented"; return nil }

func (s *basic) Close() error { _ = "STUB: not implemented"; return nil }

func (s *packageWriter) PackfileWriter() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (s *reflogBasic) Reflog(name plumbing.ReferenceName) ([]*reflog.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *reflogBasic) AppendReflog(name plumbing.ReferenceName, entry *reflog.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *reflogBasic) DeleteReflog(name plumbing.ReferenceName) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *reflogPackageWriter) PackfileWriter() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}
