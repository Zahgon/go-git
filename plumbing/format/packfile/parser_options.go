package packfile

import (
	"github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

type ParserOption func(*Parser)

func WithStorage(storage storer.EncodedObjectStorer) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithScannerObservers(ob ...Observer) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithObjectFormat(of config.ObjectFormat) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithHighMemoryMode() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }
