package protocol

import (
	"errors"
)

var ErrUnknownProtocol = errors.New("unknown Git Wire protocol")

type Version int

const (
	V0 Version = iota

	V1

	V2

	Undefined Version = -1
)

func (v Version) String() string { _ = "STUB: not implemented"; return "" }

func Parse(v string) (Version, error) { _ = "STUB: not implemented"; return *new(Version), nil }
