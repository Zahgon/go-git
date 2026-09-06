package storer

import (
	"github.com/go-git/go-billy/v6"
)

type Storer interface {
	EncodedObjectStorer
	ReferenceStorer
}

type Initializer interface {
	Init() error
}

type FilesystemStorer interface {
	Filesystem() billy.Filesystem
}

type IdleReleaser interface {
	CloseIdleDescriptors() error
}
