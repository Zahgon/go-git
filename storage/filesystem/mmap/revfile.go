//go:build darwin || linux

package mmap

import (
	"github.com/go-git/go-billy/v6"
)

var (
	revSignature = []byte{'R', 'I', 'D', 'X'}
	revMinLen    = 16
	revSupported = uint32(1)
)

const (
	revHeader = 4 + 4 + 4
)

func (s *PackScanner) loadRevFile(rev billy.File) error { _ = "STUB: not implemented"; return nil }
