package storage

import "github.com/go-git/go-git/v6/plumbing/format/config"

type ObjectFormatSetter interface {
	SetObjectFormat(config.ObjectFormat) error
}

type ExtensionChecker interface {
	SupportsExtension(name, value string) bool
}
