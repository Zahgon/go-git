package config

import "errors"

type RepositoryFormatVersion string

const (
	Version0 = "0"

	Version1 = "1"

	DefaultRepositoryFormatVersion = Version0
)

type ObjectFormat string

const (
	UnsetObjectFormat ObjectFormat = ""

	SHA1 ObjectFormat = "sha1"

	SHA256 ObjectFormat = "sha256"

	DefaultObjectFormat = SHA1
)

func (f ObjectFormat) String() string { _ = "STUB: not implemented"; return "" }

func (f ObjectFormat) Size() int { _ = "STUB: not implemented"; return 0 }

func (f ObjectFormat) HexSize() int { _ = "STUB: not implemented"; return 0 }

var ErrInvalidObjectFormat = errors.New("invalid object format")

const (
	SHA1Size = 20

	SHA256Size = 32

	SHA1HexSize = SHA1Size * 2

	SHA256HexSize = SHA256Size * 2
)
