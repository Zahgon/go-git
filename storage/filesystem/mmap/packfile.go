//go:build darwin || linux

package mmap

import (
	"github.com/go-git/go-billy/v6"
)

var (
	packSignature = []byte{'P', 'A', 'C', 'K'}
	packMinLen    = 32
	packSupported = uint32(2)
)

func (s *PackScanner) loadPackFile(pack billy.File) error { _ = "STUB: not implemented"; return nil }
