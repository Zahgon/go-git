package zlib

import (
	"errors"
)

var (
	ErrNilProvider = errors.New("zlib provider must not be nil")

	ErrNilReader = errors.New("zlib provider returned nil reader")

	ErrNilWriter = errors.New("zlib provider returned nil writer")
)

var validateInitBytes = []byte{0x78, 0x9c, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0x00, 0x01}

func ValidateProvider(provider Provider) error { _ = "STUB: not implemented"; return nil }

func isNil(v any) bool { _ = "STUB: not implemented"; return false }
